package secret

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/VladimirSh98/GophKeeper/internal/server/repository/secret"
	"github.com/VladimirSh98/GophKeeper/internal/server/repository/user"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret"
	mockUser "github.com/VladimirSh98/GophKeeper/mocks/user"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateSecret(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockUser.NewMockRepository(ctrl)
	mockSecretRepo := mockSecret.NewMockRepository(ctrl)
	logger, _ := zap.NewDevelopment()
	grpcSvc := NewSecretGrpc(mockUserRepo, mockSecretRepo, logger)

	ctx := context.WithValue(context.Background(), utils.UserLoginKey, "user1")

	t.Run("successful creation", func(t *testing.T) {
		req := &pb.CreateSecretRequest{
			DataType: pb.DataType_TEXT_DATA,
			Content:  []byte("secret content"),
			Metadata: map[string]string{"note": "test"},
		}

		userModel := user.User{ID: 1, Login: "user1"}

		mockUserRepo.EXPECT().
			GetUserByLogin(gomock.Any(), "user1", false).
			Return(userModel, nil)

		mockSecretRepo.EXPECT().
			Create(gomock.Any(), int(userModel.ID), int(req.DataType), req.Content, gomock.Any()).
			DoAndReturn(func(_ context.Context, userID, dataType int, content []byte, metadata []byte) (secret.Secret, error) {
				return secret.Secret{
					ID:        10,
					UserID:    userID,
					DataType:  secret.DataType(dataType),
					Content:   content,
					Metadata:  metadata,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil
			})

		resp, err := grpcSvc.Create(ctx, req)
		require.NoError(t, err)
		require.Equal(t, int64(10), resp.Id)
		require.Equal(t, int64(1), resp.UserId)
		require.Equal(t, []byte("secret content"), resp.Content)
		require.Equal(t, "test", resp.Metadata["note"])
		require.Equal(t, pb.DataType_TEXT_DATA, resp.DataType)
		require.WithinDuration(t, time.Now(), resp.CreatedAt.AsTime(), time.Second)
	})

	t.Run("user not found", func(t *testing.T) {
		mockUserRepo.EXPECT().
			GetUserByLogin(gomock.Any(), "user1", false).
			Return(user.User{}, errors.New("not found"))

		req := &pb.CreateSecretRequest{}
		resp, err := grpcSvc.Create(ctx, req)
		require.Nil(t, resp)
		st, _ := status.FromError(err)
		require.Equal(t, codes.PermissionDenied, st.Code())
	})

	t.Run("secret creation error", func(t *testing.T) {
		mockUserRepo.EXPECT().
			GetUserByLogin(gomock.Any(), "user1", false).
			Return(user.User{ID: 1}, nil)

		req := &pb.CreateSecretRequest{
			DataType: pb.DataType_TEXT_DATA,
			Content:  []byte("fail"),
			Metadata: map[string]string{"note": "test"},
		}

		mockSecretRepo.EXPECT().
			Create(gomock.Any(), 1, int(req.DataType), req.Content, gomock.Any()).
			Return(secret.Secret{}, errors.New("fail"))

		resp, err := grpcSvc.Create(ctx, req)
		require.Nil(t, resp)
		st, _ := status.FromError(err)
		require.Equal(t, codes.Internal, st.Code())
	})

	t.Run("metadata unmarshal error", func(t *testing.T) {
		mockUserRepo.EXPECT().
			GetUserByLogin(gomock.Any(), "user1", false).
			Return(user.User{ID: 1}, nil)

		req := &pb.CreateSecretRequest{
			DataType: pb.DataType_TEXT_DATA,
			Content:  []byte("content"),
			Metadata: map[string]string{"note": "test"},
		}

		mockSecretRepo.EXPECT().
			Create(gomock.Any(), 1, int(req.DataType), req.Content, gomock.Any()).
			Return(secret.Secret{
				ID:        1,
				UserID:    1,
				DataType:  secret.DataType(req.DataType),
				Content:   req.Content,
				Metadata:  []byte{0xff, 0xfe}, // invalid JSON
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

		resp, err := grpcSvc.Create(ctx, req)
		require.Nil(t, resp)
		st, _ := status.FromError(err)
		require.Equal(t, codes.Internal, st.Code())
	})
}

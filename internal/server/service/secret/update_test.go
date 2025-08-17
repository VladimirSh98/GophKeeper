package secret

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"

	"github.com/VladimirSh98/GophKeeper/internal/server/repository/secret"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestUpdateSecret(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSecretRepo := mockSecret.NewMockRepository(ctrl)
	logger, _ := zap.NewDevelopment()
	grpcSvc := NewSecretGrpc(nil, mockSecretRepo, logger)

	ctx := context.WithValue(context.Background(), utils.UserLoginKey, "user1")

	t.Run("successful update", func(t *testing.T) {
		metadata := map[string]string{"key": "value"}
		bytesMetadata, _ := json.Marshal(metadata)

		req := &pb.EditSecretRequest{
			Id:       1,
			Content:  []byte("new content"),
			Metadata: metadata,
		}

		mockSecretRepo.EXPECT().
			UpdateByID(ctx, "user1", int(req.Id), req.Content, bytesMetadata).
			Return(secret.Secret{
				ID:        1,
				UserID:    1,
				Archived:  false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DataType:  1,
				Content:   req.Content,
				Metadata:  bytesMetadata,
			}, nil)

		resp, err := grpcSvc.Update(ctx, req)
		require.NoError(t, err)
		require.Equal(t, int64(1), resp.Id)
		require.Equal(t, "value", resp.Metadata["key"])
	})

	t.Run("repo update error", func(t *testing.T) {
		req := &pb.EditSecretRequest{
			Id:      1,
			Content: []byte("new content"),
			Metadata: map[string]string{
				"key": "value",
			},
		}
		bytesMetadata, _ := json.Marshal(req.Metadata)

		mockSecretRepo.EXPECT().
			UpdateByID(ctx, "user1", int(req.Id), req.Content, bytesMetadata).
			Return(secret.Secret{}, errors.New("db error"))

		resp, err := grpcSvc.Update(ctx, req)
		require.Nil(t, resp)
		st, _ := status.FromError(err)
		require.Equal(t, st.Code(), codes.Internal)
		require.Contains(t, st.Message(), "secret update failed")
	})

	t.Run("metadata unmarshal error", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), utils.UserLoginKey, "user1")
		req := &pb.EditSecretRequest{
			Id:      1,
			Content: []byte("new content"),
			Metadata: map[string]string{
				"key": "value",
			},
		}
		mockSecretRepo.EXPECT().
			UpdateByID(ctx, "user1", int(req.Id), req.Content, gomock.Any()).
			Return(secret.Secret{
				ID:        1,
				UserID:    1,
				Archived:  false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DataType:  1,
				Content:   req.Content,
				Metadata:  []byte{0xff, 0xfe},
			}, nil)

		resp, err := grpcSvc.Update(ctx, req)
		require.Nil(t, resp)
		st, _ := status.FromError(err)
		require.Equal(t, codes.Internal, st.Code())
		require.Contains(t, st.Message(), "secret metadata unmarshal failed")
	})
}

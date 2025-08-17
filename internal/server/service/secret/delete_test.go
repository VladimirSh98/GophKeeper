package secret

import (
	"context"
	"errors"
	"github.com/VladimirSh98/GophKeeper/internal/server/repository/secret"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret"
	mockUser "github.com/VladimirSh98/GophKeeper/mocks/user"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestDeleteSecret(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSecretRepo := mockSecret.NewMockRepository(ctrl)
	mockUserRepo := mockUser.NewMockRepository(ctrl) // если нужен для создания Grpc
	logger, _ := zap.NewDevelopment()
	grpcSvc := NewSecretGrpc(mockUserRepo, mockSecretRepo, logger)

	ctx := context.WithValue(context.Background(), utils.UserLoginKey, "user1")

	t.Run("successful deletion", func(t *testing.T) {
		req := &pb.DeleteSecretRequest{Id: 1}

		mockSecretRepo.EXPECT().
			DeleteByID(ctx, "user1", 1).
			Return(secret.Secret{}, nil)

		resp, err := grpcSvc.Delete(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, resp)
	})

	t.Run("deletion error", func(t *testing.T) {
		req := &pb.DeleteSecretRequest{Id: 2}

		mockSecretRepo.EXPECT().
			DeleteByID(ctx, "user1", 2).
			Return(secret.Secret{}, errors.New("db error"))

		resp, err := grpcSvc.Delete(ctx, req)
		require.Nil(t, resp)
		require.Error(t, err)
		st, _ := status.FromError(err)
		require.Equal(t, codes.Internal, st.Code())
	})
}

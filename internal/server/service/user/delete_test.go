package user

import (
	"context"
	"errors"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret"
	mockUser "github.com/VladimirSh98/GophKeeper/mocks/user"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockUser.NewMockRepository(ctrl)
	mockSecretRepo := mockSecret.NewMockRepository(ctrl)
	logger, _ := zap.NewDevelopment()
	grpcSvc := NewUserGrpc(mockUserRepo, mockSecretRepo, nil, logger)

	ctx := context.WithValue(context.Background(), utils.UserLoginKey, "testuser")

	t.Run("successful delete", func(t *testing.T) {
		mockUserRepo.EXPECT().
			Delete(ctx, "testuser").
			Return(nil)
		mockSecretRepo.EXPECT().
			DeleteByLogin(ctx, "testuser").
			Return(nil)

		resp, err := grpcSvc.Delete(ctx, &pb.DeleteRequest{})
		require.NoError(t, err)
		require.NotNil(t, resp)
	})

	t.Run("user delete error", func(t *testing.T) {
		mockUserRepo.EXPECT().
			Delete(ctx, "testuser").
			Return(errors.New("db error"))

		resp, err := grpcSvc.Delete(ctx, &pb.DeleteRequest{})
		require.Nil(t, resp)
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed deleting user")
	})

	t.Run("secret delete error", func(t *testing.T) {
		mockUserRepo.EXPECT().
			Delete(ctx, "testuser").
			Return(nil)
		mockSecretRepo.EXPECT().
			DeleteByLogin(ctx, "testuser").
			Return(errors.New("db error"))

		resp, err := grpcSvc.Delete(ctx, &pb.DeleteRequest{})
		require.Nil(t, resp)
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed deleting secrets")
	})
}

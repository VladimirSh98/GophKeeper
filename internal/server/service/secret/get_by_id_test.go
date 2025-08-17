package secret

import (
	"context"
	"errors"
	"github.com/VladimirSh98/GophKeeper/internal/server/repository/secret"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	"testing"

	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret"
	mockUser "github.com/VladimirSh98/GophKeeper/mocks/user"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockUser.NewMockRepository(ctrl)
	mockSecretRepo := mockSecret.NewMockRepository(ctrl)
	logger, _ := zap.NewDevelopment()

	grpcSvc := NewSecretGrpc(mockUserRepo, mockSecretRepo, logger)

	ctx := context.WithValue(context.Background(), utils.UserLoginKey, "user1")

	t.Run("successful get", func(t *testing.T) {
		mockSecretRepo.EXPECT().
			GetSecretByIDUser(gomock.Any(), 1, "user1").
			Return(secret.Secret{
				ID:       1,
				UserID:   1,
				Content:  []byte("secret content"),
				DataType: 0,
				Metadata: []byte(`{"key":"value"}`),
			}, nil)

		resp, err := grpcSvc.GetByID(ctx, &pb.GetSecretByIDRequest{Id: 1})
		require.NoError(t, err)
		require.Equal(t, int64(1), resp.Id)
		require.Equal(t, "value", resp.Metadata["key"])
	})

	t.Run("secret repo error", func(t *testing.T) {
		mockSecretRepo.EXPECT().
			GetSecretByIDUser(ctx, 2, "user1").
			Return(secret.Secret{}, errors.New("db error"))

		resp, err := grpcSvc.GetByID(ctx, &pb.GetSecretByIDRequest{Id: 2})
		require.Nil(t, resp)
		require.Error(t, err)
	})
}

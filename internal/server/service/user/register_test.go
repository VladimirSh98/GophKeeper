package user

import (
	"context"
	"errors"
	"github.com/VladimirSh98/GophKeeper/mocks/auth"
	"github.com/VladimirSh98/GophKeeper/mocks/user"
	"testing"

	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := user.NewMockRepository(ctrl)
	mockAuth := auth.NewMockService(ctrl)
	logger, _ := zap.NewDevelopment()
	grpcSvc := NewUserGrpc(mockUserRepo, nil, mockAuth, logger)
	t.Run("successful registration", func(t *testing.T) {
		req := &pb.RegisterRequest{
			Login:    "newuser",
			Password: "password123",
		}
		mockUserRepo.EXPECT().
			Create(gomock.Any(), req.Login, gomock.Any()).
			Return(1, nil)

		mockAuth.EXPECT().
			CreateToken(req.Login).
			Return("mocktoken", nil)

		resp, err := grpcSvc.Register(context.Background(), req)
		require.NoError(t, err)
		require.Equal(t, "mocktoken", resp.Token)
	})

	t.Run("duplicate user", func(t *testing.T) {
		req := &pb.RegisterRequest{
			Login:    "existing",
			Password: "password123",
		}
		mockUserRepo.EXPECT().
			Create(gomock.Any(), req.Login, gomock.Any()).
			Return(0, errors.New("duplicate"))
		resp, err := grpcSvc.Register(context.Background(), req)
		require.Nil(t, resp)
		require.Error(t, err)
	})
}

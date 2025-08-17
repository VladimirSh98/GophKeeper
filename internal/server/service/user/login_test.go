package user

import (
	"context"
	"errors"
	"testing"

	"github.com/VladimirSh98/GophKeeper/internal/server/repository/user"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	"github.com/VladimirSh98/GophKeeper/mocks/auth"
	mockUser "github.com/VladimirSh98/GophKeeper/mocks/user"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockUser.NewMockRepository(ctrl)
	mockAuth := auth.NewMockService(ctrl)
	logger, _ := zap.NewDevelopment()
	grpcSvc := NewUserGrpc(mockUserRepo, nil, mockAuth, logger)

	t.Run("successful login", func(t *testing.T) {
		password := "password123"
		hash, err := utils.HashPassword(password)
		require.NoError(t, err)

		req := &pb.LoginRequest{
			Login:    "user1",
			Password: password,
		}

		mockUserRepo.EXPECT().
			GetUserByLogin(gomock.Any(), req.Login, false).
			Return(user.User{Login: req.Login, Hash: hash}, nil)

		mockAuth.EXPECT().
			CreateToken(req.Login).
			Return("mocktoken", nil)

		resp, err := grpcSvc.Login(context.Background(), req)
		require.NoError(t, err)
		require.Equal(t, "mocktoken", resp.Token)
	})

	t.Run("invalid password", func(t *testing.T) {
		hash, err := utils.HashPassword("correctpassword")
		require.NoError(t, err)

		req := &pb.LoginRequest{
			Login:    "user1",
			Password: "wrongpassword",
		}

		mockUserRepo.EXPECT().
			GetUserByLogin(gomock.Any(), req.Login, false).
			Return(user.User{Login: req.Login, Hash: hash}, nil)

		resp, err := grpcSvc.Login(context.Background(), req)
		require.Nil(t, resp)
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid credentials")
	})

	t.Run("db error", func(t *testing.T) {
		req := &pb.LoginRequest{
			Login:    "user1",
			Password: "password123",
		}

		mockUserRepo.EXPECT().
			GetUserByLogin(gomock.Any(), req.Login, false).
			Return(user.User{}, errors.New("db error"))

		resp, err := grpcSvc.Login(context.Background(), req)
		require.Nil(t, resp)
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to get user from db")
	})
}

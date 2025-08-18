package user

import (
	"testing"

	mockMemory "github.com/VladimirSh98/GophKeeper/mocks/token_manager"
	mockUser "github.com/VladimirSh98/GophKeeper/mocks/user_client"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userClient := mockUser.NewMockClientInterface(ctrl)
	tokenManager := mockMemory.NewMockTokenManager(ctrl)
	logger := zap.NewNop()

	svc := NewService(userClient, tokenManager, logger)
	if svc == nil {
		t.Fatal("expected non-nil service")
	}
}

func TestServiceMethods_NoPanic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userClient := mockUser.NewMockClientInterface(ctrl)
	tokenManager := mockMemory.NewMockTokenManager(ctrl)
	logger := zap.NewNop()
	svc := NewService(userClient, tokenManager, logger)
	cmd := &cobra.Command{}
	args := []string{}
	svc.Login(cmd, args)
	svc.Register(cmd, args)
	svc.Delete(cmd, args)
}

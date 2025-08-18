package user

import (
	"errors"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"testing"

	mockMemory "github.com/VladimirSh98/GophKeeper/mocks/token_manager"
	mockUser "github.com/VladimirSh98/GophKeeper/mocks/user_client"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServiceLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tokenManager := mockMemory.NewMockTokenManager(ctrl)
	userClient := mockUser.NewMockClientInterface(ctrl)

	svc := &Service{
		userClient:   userClient,
		tokenManager: tokenManager,
	}

	tests := []struct {
		name         string
		login        string
		password     string
		clientResp   *pb.LoginResponse
		clientErr    error
		saveTokenErr error
		expectSave   bool
	}{
		{
			name:         "successful login",
			login:        "user1",
			password:     "pass1",
			clientResp:   &pb.LoginResponse{Token: "token123"},
			clientErr:    nil,
			saveTokenErr: nil,
			expectSave:   true,
		},
		{
			name:         "unauthenticated error",
			login:        "user2",
			password:     "wrongpass",
			clientResp:   nil,
			clientErr:    status.Error(codes.Unauthenticated, "unauthenticated"),
			saveTokenErr: nil,
			expectSave:   false,
		},
		{
			name:         "unexpected error",
			login:        "user3",
			password:     "pass3",
			clientResp:   nil,
			clientErr:    errors.New("some error"),
			saveTokenErr: nil,
			expectSave:   false,
		},
		{
			name:         "save token error",
			login:        "user4",
			password:     "pass4",
			clientResp:   &pb.LoginResponse{Token: "token456"},
			clientErr:    nil,
			saveTokenErr: errors.New("save error"),
			expectSave:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().String("login", "", "")
			cmd.Flags().String("password", "", "")
			cmd.Flags().Set("login", tt.login)
			cmd.Flags().Set("password", tt.password)

			userClient.EXPECT().Login(gomock.Any(), tt.login, tt.password).Return(tt.clientResp, tt.clientErr)
			if tt.expectSave && tt.clientResp != nil && tt.clientErr == nil {
				tokenManager.EXPECT().SaveToken(tt.clientResp.Token).Return(tt.saveTokenErr)
			}

			svc.Login(cmd, []string{})
		})
	}
}

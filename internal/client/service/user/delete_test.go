package user

import (
	"errors"
	"testing"

	mockMemory "github.com/VladimirSh98/GophKeeper/mocks/token_manager"
	mockUser "github.com/VladimirSh98/GophKeeper/mocks/user_client"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
)

func TestServiceDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tokenManager := mockMemory.NewMockTokenManager(ctrl)
	userClient := mockUser.NewMockClientInterface(ctrl)

	svc := &Service{
		userClient:   userClient,
		tokenManager: tokenManager,
	}

	tests := []struct {
		name             string
		tokenReturn      string
		tokenErr         error
		deleteErr        error
		clearTokenErr    error
		expectDeleteCall bool
		expectClearCall  bool
	}{
		{
			name:             "successful deletion",
			tokenReturn:      "token123",
			tokenErr:         nil,
			deleteErr:        nil,
			clearTokenErr:    nil,
			expectDeleteCall: true,
			expectClearCall:  true,
		},
		{
			name:             "token manager error",
			tokenReturn:      "",
			tokenErr:         errors.New("fail"),
			deleteErr:        nil,
			clearTokenErr:    nil,
			expectDeleteCall: false,
			expectClearCall:  false,
		},
		{
			name:             "empty token",
			tokenReturn:      "",
			tokenErr:         nil,
			deleteErr:        nil,
			clearTokenErr:    nil,
			expectDeleteCall: false,
			expectClearCall:  false,
		},
		{
			name:             "delete fails",
			tokenReturn:      "token123",
			tokenErr:         nil,
			deleteErr:        errors.New("delete fail"),
			clearTokenErr:    nil,
			expectDeleteCall: true,
			expectClearCall:  false,
		},
		{
			name:             "clear token fails",
			tokenReturn:      "token123",
			tokenErr:         nil,
			deleteErr:        nil,
			clearTokenErr:    errors.New("clear fail"),
			expectDeleteCall: true,
			expectClearCall:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}

			tokenManager.EXPECT().GetToken().Return(tt.tokenReturn, tt.tokenErr)
			if tt.expectDeleteCall {
				userClient.EXPECT().Delete(gomock.Any(), tt.tokenReturn).Return(nil, tt.deleteErr)
			}
			if tt.expectClearCall {
				tokenManager.EXPECT().ClearToken().Return(tt.clearTokenErr)
			}

			svc.Delete(cmd, []string{})
		})
	}
}

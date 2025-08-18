package secret

import (
	"errors"
	"testing"

	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret_client"
	mockToken "github.com/VladimirSh98/GophKeeper/mocks/token_manager"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
)

func TestServiceDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tokenManager := mockToken.NewMockTokenManager(ctrl)
	secretClient := mockSecret.NewMockClientInterface(ctrl)

	s := &Service{
		tokenManager: tokenManager,
		secretClient: secretClient,
	}

	tests := []struct {
		name             string
		tokenReturn      string
		tokenErr         error
		secretID         int
		secretErr        error
		expectSecretCall bool
	}{
		{
			name:             "success",
			tokenReturn:      "fake-token",
			tokenErr:         nil,
			secretID:         1,
			secretErr:        nil,
			expectSecretCall: true,
		},
		{
			name:             "token manager error",
			tokenReturn:      "",
			tokenErr:         errors.New("fail"),
			secretID:         1,
			secretErr:        nil,
			expectSecretCall: false,
		},
		{
			name:             "empty token",
			tokenReturn:      "",
			tokenErr:         nil,
			secretID:         1,
			secretErr:        nil,
			expectSecretCall: false,
		},
		{
			name:             "secret client error",
			tokenReturn:      "fake-token",
			tokenErr:         nil,
			secretID:         1,
			secretErr:        errors.New("delete fail"),
			expectSecretCall: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().Int("secret", tt.secretID, "")

			tokenManager.EXPECT().GetToken().Return(tt.tokenReturn, tt.tokenErr)

			if tt.expectSecretCall {
				secretClient.EXPECT().
					Delete(gomock.Any(), tt.tokenReturn, tt.secretID).
					Return(&pb.SecretModel{Id: int64(tt.secretID)}, tt.secretErr)
			}

			s.Delete(cmd, []string{})
		})
	}
}

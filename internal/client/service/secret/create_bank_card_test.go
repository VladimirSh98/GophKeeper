package secret

import (
	"errors"
	"github.com/VladimirSh98/GophKeeper/internal/client/repository/secret"
	"testing"

	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret_client"
	mockToken "github.com/VladimirSh98/GophKeeper/mocks/token_manager"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
)

func TestServiceCreateBankCard(t *testing.T) {
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
		secretErr        error
		expectSecretCall bool
	}{
		{
			name:             "success",
			tokenReturn:      "fake-token",
			tokenErr:         nil,
			secretErr:        nil,
			expectSecretCall: true,
		},
		{
			name:             "token manager error",
			tokenReturn:      "",
			tokenErr:         errors.New("fail"),
			secretErr:        nil,
			expectSecretCall: false,
		},
		{
			name:             "empty token",
			tokenReturn:      "",
			tokenErr:         nil,
			secretErr:        nil,
			expectSecretCall: false,
		},
		{
			name:             "secret client error",
			tokenReturn:      "fake-token",
			tokenErr:         nil,
			secretErr:        errors.New("secret fail"),
			expectSecretCall: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().String("number", "1234567890123456", "")
			cmd.Flags().String("expiry_month", "12", "")
			cmd.Flags().String("expiry_year", "25", "")
			cmd.Flags().String("holder_name", "John Doe", "")
			cmd.Flags().String("cvv", "123", "")
			cmd.Flags().StringArray("metadata", []string{"key=value"}, "")

			tokenManager.EXPECT().GetToken().Return(tt.tokenReturn, tt.tokenErr)

			if tt.expectSecretCall {
				secretClient.EXPECT().
					Create(gomock.Any(), tt.tokenReturn, secret.BankCard, gomock.Any(), map[string]string{"key": "value"}).
					Return(&pb.SecretModel{Id: 1}, tt.secretErr)
			}

			s.CreateBankCard(cmd, []string{})
		})
	}
}

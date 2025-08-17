package secret

import (
	"errors"
	"google.golang.org/protobuf/proto"
	"testing"

	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret_client"
	mockToken "github.com/VladimirSh98/GophKeeper/mocks/token_manager"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
)

func TestServiceGet(t *testing.T) {
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
		clientResp       *pb.GetSecretsResponse
		clientErr        error
		expectClientCall bool
	}{
		{
			name:             "success",
			tokenReturn:      "fake-token",
			tokenErr:         nil,
			clientResp:       &pb.GetSecretsResponse{Secrets: []*pb.SecretModel{{Id: 1, DataType: pb.DataType_TEXT_DATA, Content: []byte("data")}}},
			clientErr:        nil,
			expectClientCall: true,
		},
		{
			name:             "token manager error",
			tokenReturn:      "",
			tokenErr:         errors.New("fail"),
			clientResp:       nil,
			clientErr:        nil,
			expectClientCall: false,
		},
		{
			name:             "empty token",
			tokenReturn:      "",
			tokenErr:         nil,
			clientResp:       nil,
			clientErr:        nil,
			expectClientCall: false,
		},
		{
			name:             "client error",
			tokenReturn:      "fake-token",
			tokenErr:         nil,
			clientResp:       nil,
			clientErr:        errors.New("client fail"),
			expectClientCall: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}

			tokenManager.EXPECT().GetToken().Return(tt.tokenReturn, tt.tokenErr)

			if tt.expectClientCall {
				secretClient.EXPECT().Get(gomock.Any(), tt.tokenReturn).Return(tt.clientResp, tt.clientErr)
			}

			s.Get(cmd, []string{})
		})
	}
}

func TestPrintResponse(t *testing.T) {
	loginPassContent, _ := proto.Marshal(&pb.LoginPass{Login: "user", Password: "pass"})
	bankCardContent, _ := proto.Marshal(&pb.BankCard{Number: "1234", HolderName: "John", ExpiryMonth: "01", ExpiryYear: "25", Cvv: "123"})
	textContent, _ := proto.Marshal(&pb.TextData{Text: "Some text"})
	binaryContent, _ := proto.Marshal(&pb.BinaryData{Data: []byte{0x01, 0x02, 0x03}})

	response := &pb.GetSecretsResponse{
		Secrets: []*pb.SecretModel{
			{Id: 1, DataType: pb.DataType_LOGIN_PASSWORD, Content: loginPassContent},
			{Id: 2, DataType: pb.DataType_BANK_CARD, Content: bankCardContent},
			{Id: 3, DataType: pb.DataType_TEXT_DATA, Content: textContent},
			{Id: 4, DataType: pb.DataType_BINARY_DATA, Content: binaryContent},
			{Id: 5, DataType: 999, Content: []byte("invalid")},
		},
	}
	printResponse(response)
}

package secret

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"

	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret_client"
	mockToken "github.com/VladimirSh98/GophKeeper/mocks/token_manager"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

func mustMarshal(m proto.Message) []byte {
	b, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	return b
}

func TestServiceUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tokenManager := mockToken.NewMockTokenManager(ctrl)
	secretClient := mockSecret.NewMockClientInterface(ctrl)

	s := &Service{
		tokenManager: tokenManager,
		secretClient: secretClient,
	}

	oldContent, _ := proto.Marshal(&pb.TextData{Text: "old"})

	tests := []struct {
		name             string
		tokenReturn      string
		tokenErr         error
		secretResp       *pb.SecretModel
		secretErr        error
		expectClientCall bool
	}{
		{
			name:        "success update",
			tokenReturn: "token",
			tokenErr:    nil,
			secretResp: &pb.SecretModel{
				Id:       1,
				DataType: pb.DataType_TEXT_DATA,
				Content:  oldContent,
				Metadata: map[string]string{"k": "v"},
			},
			secretErr:        nil,
			expectClientCall: true,
		},
		{
			name:             "token manager error",
			tokenReturn:      "",
			tokenErr:         errors.New("fail"),
			secretResp:       nil,
			secretErr:        nil,
			expectClientCall: false,
		},
		{
			name:             "empty token",
			tokenReturn:      "",
			tokenErr:         nil,
			secretResp:       nil,
			secretErr:        nil,
			expectClientCall: false,
		},
		{
			name:             "get secret error",
			tokenReturn:      "token",
			tokenErr:         nil,
			secretResp:       nil,
			secretErr:        errors.New("get secret fail"),
			expectClientCall: false,
		},
		{
			name:        "update error",
			tokenReturn: "token",
			tokenErr:    nil,
			secretResp: &pb.SecretModel{
				Id:       1,
				DataType: pb.DataType_TEXT_DATA,
				Content:  oldContent,
			},
			secretErr:        nil,
			expectClientCall: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().Int("secret", 1, "")
			cmd.Flags().String("text", "new", "")

			tokenManager.EXPECT().GetToken().Return(tt.tokenReturn, tt.tokenErr)

			if tt.secretResp != nil || tt.secretErr != nil {
				secretClient.EXPECT().GetByID(gomock.Any(), tt.tokenReturn, 1).Return(tt.secretResp, tt.secretErr)
			}

			if tt.expectClientCall {
				secretClient.EXPECT().Update(gomock.Any(), tt.tokenReturn, 1, gomock.Any(), gomock.Any()).Return(tt.secretResp, nil)
			}

			s.Update(cmd, []string{})
		})
	}
}

func TestGetContentBySecret(t *testing.T) {
	tests := []struct {
		name   string
		secret *pb.SecretModel
		flags  map[string]string
		check  func(t *testing.T, out []byte, secret *pb.SecretModel)
	}{
		{
			name: "parseSecret error (unknown DataType)",
			secret: &pb.SecretModel{
				DataType: pb.DataType(999),
				Content:  []byte("raw"),
			},
			flags: map[string]string{},
			check: func(t *testing.T, out []byte, secret *pb.SecretModel) {
				require.Equal(t, secret.Content, out)
			},
		},
		{
			name: "login_password unchanged",
			secret: &pb.SecretModel{
				DataType: pb.DataType_LOGIN_PASSWORD,
				Content:  mustMarshal(&pb.LoginPass{Login: "user", Password: "pass"}),
			},
			flags: map[string]string{},
			check: func(t *testing.T, out []byte, _ *pb.SecretModel) {
				lp := &pb.LoginPass{}
				require.NoError(t, proto.Unmarshal(out, lp))
				require.Equal(t, "user", lp.Login)
				require.Equal(t, "pass", lp.Password)
			},
		},
		{
			name: "login_password updated",
			secret: &pb.SecretModel{
				DataType: pb.DataType_LOGIN_PASSWORD,
				Content:  mustMarshal(&pb.LoginPass{Login: "old", Password: "old"}),
			},
			flags: map[string]string{"login": "newuser", "password": "newpass"},
			check: func(t *testing.T, out []byte, _ *pb.SecretModel) {
				lp := &pb.LoginPass{}
				require.NoError(t, proto.Unmarshal(out, lp))
				require.Equal(t, "newuser", lp.Login)
				require.Equal(t, "newpass", lp.Password)
			},
		},
		{
			name: "bank_card unchanged",
			secret: &pb.SecretModel{
				DataType: pb.DataType_BANK_CARD,
				Content: mustMarshal(&pb.BankCard{
					Number: "1111", ExpiryMonth: "12", ExpiryYear: "25", Cvv: "123",
				}),
			},
			flags: map[string]string{},
			check: func(t *testing.T, out []byte, _ *pb.SecretModel) {
				card := &pb.BankCard{}
				require.NoError(t, proto.Unmarshal(out, card))
				require.Equal(t, "1111", card.Number)
			},
		},
		{
			name: "bank_card updated",
			secret: &pb.SecretModel{
				DataType: pb.DataType_BANK_CARD,
				Content: mustMarshal(&pb.BankCard{
					Number: "1111", ExpiryMonth: "12", ExpiryYear: "25", Cvv: "123",
				}),
			},
			flags: map[string]string{"number": "2222", "cvc": "999"},
			check: func(t *testing.T, out []byte, _ *pb.SecretModel) {
				card := &pb.BankCard{}
				require.NoError(t, proto.Unmarshal(out, card))
				require.Equal(t, "2222", card.Number)
				require.Equal(t, "123", card.Cvv)
			},
		},
		{
			name: "text_data unchanged",
			secret: &pb.SecretModel{
				DataType: pb.DataType_TEXT_DATA,
				Content:  mustMarshal(&pb.TextData{Text: "test"}),
			},
			flags: map[string]string{},
			check: func(t *testing.T, out []byte, _ *pb.SecretModel) {
				td := &pb.TextData{}
				require.NoError(t, proto.Unmarshal(out, td))
				require.Equal(t, "test", td.Text)
			},
		},
		{
			name: "text_data updated",
			secret: &pb.SecretModel{
				DataType: pb.DataType_TEXT_DATA,
				Content:  mustMarshal(&pb.TextData{Text: "old1"}),
			},
			flags: map[string]string{"text": "new text"},
			check: func(t *testing.T, out []byte, _ *pb.SecretModel) {
				td := &pb.TextData{}
				require.NoError(t, proto.Unmarshal(out, td))
				require.Equal(t, "new text", td.Text)
			},
		},
		{
			name: "binary_data unchanged",
			secret: &pb.SecretModel{
				DataType: pb.DataType_BINARY_DATA,
				Content:  mustMarshal(&pb.BinaryData{Data: []byte("bin")}),
			},
			flags: map[string]string{},
			check: func(t *testing.T, out []byte, _ *pb.SecretModel) {
				bd := &pb.BinaryData{}
				require.NoError(t, proto.Unmarshal(out, bd))
				require.Equal(t, []byte("bin"), bd.Data)
			},
		},
		{
			name: "binary_data updated",
			secret: &pb.SecretModel{
				DataType: pb.DataType_BINARY_DATA,
				Content:  mustMarshal(&pb.BinaryData{Data: []byte("oldbin")}),
			},
			flags: map[string]string{"binary": "newbin"},
			check: func(t *testing.T, out []byte, _ *pb.SecretModel) {
				bd := &pb.BinaryData{}
				require.NoError(t, proto.Unmarshal(out, bd))
				require.Equal(t, []byte("newbin"), bd.Data)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			for k, v := range tt.flags {
				cmd.Flags().String(k, v, "")
			}
			out := getContentBySecret(cmd, tt.secret)
			tt.check(t, out, tt.secret)
		})
	}
}

func TestPrepareLoginPass(t *testing.T) {
	tests := []struct {
		name        string
		cmdFlags    map[string]string
		contentData proto.Message
		oldContent  []byte
		wantLogin   string
		wantPass    string
	}{
		{
			name:     "no flags, keep old",
			cmdFlags: map[string]string{},
			contentData: &pb.LoginPass{
				Login:    "oldLogin",
				Password: "oldPass",
			},
			oldContent: nil,
			wantLogin:  "oldLogin",
			wantPass:   "oldPass",
		},
		{
			name: "update login only",
			cmdFlags: map[string]string{
				"login": "newLogin",
			},
			contentData: &pb.LoginPass{
				Login:    "oldLogin",
				Password: "oldPass",
			},
			oldContent: nil,
			wantLogin:  "newLogin",
			wantPass:   "oldPass",
		},
		{
			name: "update password only",
			cmdFlags: map[string]string{
				"password": "newPass",
			},
			contentData: &pb.LoginPass{
				Login:    "oldLogin",
				Password: "oldPass",
			},
			oldContent: nil,
			wantLogin:  "oldLogin",
			wantPass:   "newPass",
		},
		{
			name: "update both",
			cmdFlags: map[string]string{
				"login":    "newLogin",
				"password": "newPass",
			},
			contentData: &pb.LoginPass{
				Login:    "oldLogin",
				Password: "oldPass",
			},
			oldContent: nil,
			wantLogin:  "newLogin",
			wantPass:   "newPass",
		},
		{
			name:        "wrong type returns old",
			cmdFlags:    map[string]string{},
			contentData: &pb.TextData{Text: "not loginpass"},
			oldContent:  []byte("oldContent"),
			wantLogin:   "",
			wantPass:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			for k, v := range tt.cmdFlags {
				cmd.Flags().String(k, v, "")
			}

			result := prepareLoginPass(cmd, tt.contentData, tt.oldContent)

			if _, ok := tt.contentData.(*pb.LoginPass); ok {
				lp := &pb.LoginPass{}
				err := proto.Unmarshal(result, lp)
				require.NoError(t, err)
				require.Equal(t, tt.wantLogin, lp.Login)
				require.Equal(t, tt.wantPass, lp.Password)
			} else {
				require.Equal(t, tt.oldContent, result)
			}
		})
	}
}

func TestPrepareBankCard(t *testing.T) {
	tests := []struct {
		name        string
		cmdFlags    map[string]string
		contentData proto.Message
		oldContent  []byte
		wantCard    *pb.BankCard
	}{
		{
			name:     "no flags, keep old",
			cmdFlags: map[string]string{},
			contentData: &pb.BankCard{
				Number:      "1111",
				HolderName:  "Old Holder",
				ExpiryMonth: "01",
				ExpiryYear:  "2030",
				Cvv:         "123",
			},
			oldContent: nil,
			wantCard: &pb.BankCard{
				Number:      "1111",
				HolderName:  "Old Holder",
				ExpiryMonth: "01",
				ExpiryYear:  "2030",
				Cvv:         "123",
			},
		},
		{
			name: "update some fields",
			cmdFlags: map[string]string{
				"number":      "2222",
				"holder_name": "New Holder",
			},
			contentData: &pb.BankCard{
				Number:      "1111",
				HolderName:  "Old Holder",
				ExpiryMonth: "01",
				ExpiryYear:  "2030",
				Cvv:         "123",
			},
			oldContent: nil,
			wantCard: &pb.BankCard{
				Number:      "2222",
				HolderName:  "New Holder",
				ExpiryMonth: "01",
				ExpiryYear:  "2030",
				Cvv:         "123",
			},
		},
		{
			name: "wrong type returns old",
			cmdFlags: map[string]string{
				"number": "9999",
			},
			contentData: &pb.TextData{Text: "not a card"},
			oldContent:  []byte("oldContent"),
			wantCard:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			for k, v := range tt.cmdFlags {
				cmd.Flags().String(k, v, "")
			}

			result := prepareBankCard(cmd, tt.contentData, tt.oldContent)

			if _, ok := tt.contentData.(*pb.BankCard); ok {
				bc := &pb.BankCard{}
				err := proto.Unmarshal(result, bc)
				require.NoError(t, err)
				require.Equal(t, tt.wantCard.Number, bc.Number)
				require.Equal(t, tt.wantCard.HolderName, bc.HolderName)
				require.Equal(t, tt.wantCard.ExpiryMonth, bc.ExpiryMonth)
				require.Equal(t, tt.wantCard.ExpiryYear, bc.ExpiryYear)
				require.Equal(t, tt.wantCard.Cvv, bc.Cvv)
			} else {
				require.Equal(t, tt.oldContent, result)
			}
		})
	}
}

func TestPrepareText(t *testing.T) {
	tests := []struct {
		name        string
		cmdFlags    map[string]string
		contentData proto.Message
		oldContent  []byte
		wantText    string
	}{
		{
			name:        "no flags, keep old text",
			cmdFlags:    map[string]string{},
			contentData: &pb.TextData{Text: "old text"},
			oldContent:  nil,
			wantText:    "old text",
		},
		{
			name: "update text via flag",
			cmdFlags: map[string]string{
				"text": "new text",
			},
			contentData: &pb.TextData{Text: "old text"},
			oldContent:  nil,
			wantText:    "new text",
		},
		{
			name:        "wrong type returns oldContent",
			cmdFlags:    map[string]string{"text": "new text"},
			contentData: &pb.LoginPass{Login: "user", Password: "pass"},
			oldContent:  []byte("old content"),
			wantText:    "old content",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			for k, v := range tt.cmdFlags {
				cmd.Flags().String(k, v, "")
			}

			result := prepareText(cmd, tt.contentData, tt.oldContent)

			if _, ok := tt.contentData.(*pb.TextData); ok {
				td := &pb.TextData{}
				err := proto.Unmarshal(result, td)
				require.NoError(t, err)
				require.Equal(t, tt.wantText, td.Text)
			} else {
				require.Equal(t, tt.oldContent, result)
			}
		})
	}
}

func TestPrepareBinary(t *testing.T) {
	tests := []struct {
		name        string
		cmdFlags    map[string]string
		contentData proto.Message
		oldContent  []byte
		wantData    []byte
	}{
		{
			name:        "no binary flag, keep old data",
			cmdFlags:    map[string]string{},
			contentData: &pb.BinaryData{Data: []byte("old data")},
			oldContent:  nil,
			wantData:    []byte("old data"),
		},
		{
			name: "update binary via flag",
			cmdFlags: map[string]string{
				"binary": "new data",
			},
			contentData: &pb.BinaryData{Data: []byte("old data")},
			oldContent:  nil,
			wantData:    []byte("new data"),
		},
		{
			name:        "wrong type returns oldContent",
			cmdFlags:    map[string]string{"binary": "new data"},
			contentData: &pb.TextData{Text: "text"},
			oldContent:  []byte("old content"),
			wantData:    []byte("old content"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			for k, v := range tt.cmdFlags {
				cmd.Flags().String(k, v, "")
			}

			result := prepareBinary(cmd, tt.contentData, tt.oldContent)

			if _, ok := tt.contentData.(*pb.BinaryData); ok {
				bd := &pb.BinaryData{}
				err := proto.Unmarshal(result, bd)
				require.NoError(t, err)
				require.Equal(t, tt.wantData, bd.Data)
			} else {
				require.Equal(t, tt.oldContent, result)
			}
		})
	}
}

func TestGetMetadataBySecret(t *testing.T) {
	tests := []struct {
		name       string
		flags      []string
		secret     *pb.SecretModel
		expectedMD map[string]string
	}{
		{
			name:  "metadata from flags",
			flags: []string{"key1=value1", "key2=value2"},
			secret: &pb.SecretModel{
				Metadata: map[string]string{"old": "data"},
			},
			expectedMD: map[string]string{"key1": "value1", "key2": "value2"},
		},
		{
			name:  "no flags, use secret metadata",
			flags: []string{},
			secret: &pb.SecretModel{
				Metadata: map[string]string{"old": "data"},
			},
			expectedMD: map[string]string{"old": "data"},
		},
		{
			name:  "invalid flag ignored",
			flags: []string{"invalidflag"},
			secret: &pb.SecretModel{
				Metadata: map[string]string{"old": "data"},
			},
			expectedMD: map[string]string{"old": "data"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().StringArray("metadata", tt.flags, "")

			result := getMetadataBySecret(cmd, tt.secret)
			require.Equal(t, tt.expectedMD, result)
		})
	}
}

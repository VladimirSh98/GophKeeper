package secret

import (
	"testing"

	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestParseSecret(t *testing.T) {
	tests := []struct {
		name        string
		contentType pb.DataType
		content     proto.Message
		wantErr     bool
	}{
		{
			name:        "login password success",
			contentType: pb.DataType_LOGIN_PASSWORD,
			content:     &pb.LoginPass{Login: "user", Password: "pass"},
			wantErr:     false,
		},
		{
			name:        "bank card success",
			contentType: pb.DataType_BANK_CARD,
			content:     &pb.BankCard{Number: "1234", Cvv: "123"},
			wantErr:     false,
		},
		{
			name:        "text data success",
			contentType: pb.DataType_TEXT_DATA,
			content:     &pb.TextData{Text: "hello"},
			wantErr:     false,
		},
		{
			name:        "binary data success",
			contentType: pb.DataType_BINARY_DATA,
			content:     &pb.BinaryData{Data: []byte{1, 2, 3}},
			wantErr:     false,
		},
		{
			name:        "unknown type",
			contentType: 999,
			content:     nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var data []byte
			if tt.content != nil {
				data, _ = proto.Marshal(tt.content)
			}
			res, err := parseSecret(tt.contentType, data)
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, res)
			} else {
				require.NoError(t, err)
				require.NotNil(t, res)
			}
		})
	}
}

func TestFormatMetadata(t *testing.T) {
	tests := []struct {
		name string
		md   map[string]string
		want string
	}{
		{
			name: "normal metadata",
			md:   map[string]string{"key1": "value1", "key2": "value2"},
			want: "key1=value1 key2=value2",
		},
		{
			name: "empty metadata",
			md:   map[string]string{},
			want: "(нет метаданных)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := formatMetadata(tt.md)
			// Поскольку порядок в map не гарантирован, проверим наличие подстрок
			if len(tt.md) > 0 {
				for k, v := range tt.md {
					require.Contains(t, res, k+"="+v)
				}
			} else {
				require.Equal(t, tt.want, res)
			}
		})
	}
}

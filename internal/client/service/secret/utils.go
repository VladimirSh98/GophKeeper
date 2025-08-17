package secret

import (
	"fmt"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"google.golang.org/protobuf/proto"
	"strings"
)

func parseSecret(contentType pb.DataType, content []byte) (proto.Message, error) {
	switch contentType {
	case pb.DataType_LOGIN_PASSWORD:
		var data pb.LoginPass
		if err := proto.Unmarshal(content, &data); err != nil {
			return nil, fmt.Errorf("ошибка распаковки LoginPassword: %w", err)
		}
		return &data, nil

	case pb.DataType_BANK_CARD:
		var data pb.BankCard
		if err := proto.Unmarshal(content, &data); err != nil {
			return nil, fmt.Errorf("ошибка распаковки BankCard: %w", err)
		}
		return &data, nil

	case pb.DataType_TEXT_DATA:
		var data pb.TextData
		if err := proto.Unmarshal(content, &data); err != nil {
			return nil, fmt.Errorf("ошибка распаковки TextData: %w", err)
		}
		return &data, nil

	case pb.DataType_BINARY_DATA:
		var data pb.BinaryData
		if err := proto.Unmarshal(content, &data); err != nil {
			return nil, fmt.Errorf("ошибка распаковки BinaryData: %w", err)
		}
		return &data, nil

	default:
		return nil, fmt.Errorf("неизвестный тип секрета: %v", contentType)
	}
}

func formatMetadata(md map[string]string) string {
	if len(md) == 0 {
		return "(нет метаданных)"
	}
	var parts []string
	for k, v := range md {
		parts = append(parts, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(parts, " ")
}

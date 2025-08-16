package secret

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/client/repository/secret"
	"github.com/VladimirSh98/GophKeeper/internal/client/utils"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
	"strings"
)

// CreateBinary create binary secret
func (s *Service) CreateBinary(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	token, err := s.tokenManager.GetToken()
	if err != nil {
		utils.ColorMessage("Произошла непредвиденная ошибка", color.FgHiRed)
		return
	}
	if token == "" {
		utils.ColorMessage("Для удаления аккаунта необходимо авторизоваться", color.FgHiYellow)
		return
	}
	binary, _ := cmd.Flags().GetString("binary")
	binaryData := []byte(binary)
	metadataArgs, _ := cmd.Flags().GetStringArray("metadata")
	metadata := make(map[string]string)
	for _, kv := range metadataArgs {
		parts := strings.SplitN(kv, "=", 2)
		if len(parts) == 2 {
			metadata[parts[0]] = parts[1]
		}
	}
	var content []byte
	content, err = proto.Marshal(&pb.BinaryData{Data: binaryData})
	if err != nil {
		utils.ColorMessage("Произошла непредвиденная ошибка", color.FgHiRed)
		return
	}
	_, err = s.secretClient.Create(ctx, token, secret.BinaryData, content, metadata)
	if err != nil {
		utils.ColorMessage("Произошла непредвиденная ошибка", color.FgHiRed)
		return
	}
	utils.ColorMessage("Секрет успешно сохранен", color.FgHiGreen)
}

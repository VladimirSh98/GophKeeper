package secret

import (
	"context"
	"fmt"
	"github.com/VladimirSh98/GophKeeper/internal/client/utils"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

// Get all secrets
func (s *Service) Get(cmd *cobra.Command, args []string) {
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
	var response *pb.GetSecretsResponse
	response, err = s.secretClient.Get(ctx, token)
	if err != nil {
		utils.ColorMessage("Произошла непредвиденная ошибка", color.FgHiRed)
		return
	}
	utils.ColorMessage("Твои секретики:", color.FgHiGreen)
	printResponse(response)
}

func printResponse(response *pb.GetSecretsResponse) {
	for _, secret := range response.GetSecrets() {
		record, err := parseSecret(secret.DataType, secret.Content)
		if err != nil {
			continue
		}
		var out []byte
		out, err = protojson.MarshalOptions{
			Multiline: true,
			Indent:    "  ",
		}.Marshal(record)
		if err != nil {
			continue
		}
		utils.ColorMessage("Номер секрета: %s", color.FgHiBlue, fmt.Sprintf("%d", secret.Id))
		utils.ColorMessage("Секрет: %s", color.FgHiBlue, string(out))
		if len(secret.Metadata) > 0 {
			utils.ColorMessage("Дополнительные данные секрета: %s", color.FgHiBlue, secret.Metadata)
		}
		utils.ColorMessage("---", color.FgHiWhite)
	}
}

package secret

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/client/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Delete secret
func (s *Service) Delete(cmd *cobra.Command, args []string) {
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
	secretID, _ := cmd.Flags().GetInt64("secret")
	_, err = s.secretClient.Delete(ctx, token, int(secretID))
	if err != nil {
		utils.ColorMessage("Произошла непредвиденная ошибка", color.FgHiRed)
		return
	}
}

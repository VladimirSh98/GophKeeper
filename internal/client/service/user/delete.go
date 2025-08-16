package user

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/client/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func (s *Service) Delete(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	token, err := s.tokenManager.GetToken()
	if err != nil {
		utils.MessageWithoutArgs("Произошла непредвиденная ошибка", color.FgHiRed)
		return
	}
	if token == "" {
		utils.MessageWithoutArgs("Для удаления аккаунта необходимо авторизоваться", color.FgHiYellow)
		return
	}
	_, err = s.userClient.Delete(ctx, token)
	if err != nil {
		utils.MessageWithoutArgs("Произошла непредвиденная ошибка", color.FgHiRed)
		return
	}
	utils.MessageWithArgs("Аккаунт успешно удален!", color.FgHiGreen)
	err = s.tokenManager.ClearToken()
	if err != nil {
		return
	}
}

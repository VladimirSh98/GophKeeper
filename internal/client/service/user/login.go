package user

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/client/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
)

func (s *Service) Login(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	login, _ := cmd.Flags().GetString("login")
	password, _ := cmd.Flags().GetString("password")
	response, err := s.userClient.Login(ctx, login, password)
	if utils.IsExpectedError(err, []codes.Code{codes.Unauthenticated}) {
		utils.MessageWithArgs("Ошибка авторизации для пользователя %s", color.FgHiYellow, login)
		return
	} else if err != nil {
		utils.MessageWithoutArgs("Произошла непредвиденная ошибка", color.FgHiRed)
		return
	}
	err = s.tokenManager.SaveToken(response.Token)
	if err != nil {
		utils.MessageWithoutArgs("Произошла непредвиденная ошибка", color.FgHiRed)
		return
	}
	utils.MessageWithArgs("Пользователь %s успешно зарегистрирован!", color.FgHiGreen, login)
}

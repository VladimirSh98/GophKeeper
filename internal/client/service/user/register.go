package user

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/client/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
)

// Register register user
func (s *Service) Register(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	login, _ := cmd.Flags().GetString("login")
	password, _ := cmd.Flags().GetString("password")
	response, err := s.userClient.Register(ctx, login, password)
	if utils.IsExpectedError(err, []codes.Code{codes.FailedPrecondition}) {
		utils.ColorMessage("Пользователь с таким именем уже существует", color.FgHiYellow)
		return
	} else if err != nil {
		utils.ColorMessage("Произошла непредвиденная ошибка", color.FgHiRed)
		return
	}
	err = s.tokenManager.SaveToken(response.Token)
	if err != nil {
		utils.ColorMessage("Произошла непредвиденная ошибка", color.FgHiRed)
		return
	}
	utils.ColorMessage("Пользователь с логином %s успешно создан!", color.FgHiGreen, login)
}

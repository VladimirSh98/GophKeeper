package user

import "github.com/spf13/cobra"

// LoginCmd login user command
func (h *Handler) LoginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Авторизация пользователя",
		Run:   h.userService.Login,
	}
	cmd.Flags().StringP("login", "l", "", "Логин пользователя (обязательно)")
	cmd.Flags().StringP("password", "p", "", "Пароль пользователя (обязательно)")
	cmd.MarkFlagRequired("login")
	cmd.MarkFlagRequired("password")
	return cmd
}

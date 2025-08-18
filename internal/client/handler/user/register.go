package user

import "github.com/spf13/cobra"

// RegisterCmd register user command
func (h *Handler) RegisterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Регистрация пользователя",
		Run:   h.userService.Register,
	}
	cmd.Flags().StringP("login", "l", "", "Логин пользователя (обязательно)")
	cmd.Flags().StringP("password", "p", "", "Пароль пользователя (обязательно)")
	cmd.MarkFlagRequired("login")
	cmd.MarkFlagRequired("password")
	return cmd
}

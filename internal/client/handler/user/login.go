package user

import "github.com/spf13/cobra"

// LoginCmd login user
func (h *Handler) LoginCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Авторизация пользователя",
		Run:   h.userService.Login,
	}
}

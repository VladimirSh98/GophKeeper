package user

import "github.com/spf13/cobra"

// RegisterCmd register user
func (h *Handler) RegisterCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "register",
		Short: "Регистрация пользователя",
		Run:   h.userService.Register,
	}
}

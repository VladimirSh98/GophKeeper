package user

import "github.com/spf13/cobra"

// DeleteCmd delete user
func (h *Handler) DeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete user",
		Short: "Удаление пользователя",
		Run: func(cmd *cobra.Command, args []string) {
			println("Пользователь удален!")
		},
	}
}

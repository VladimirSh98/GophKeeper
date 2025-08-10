package user

import "github.com/spf13/cobra"

// LoginCmd login user
func LoginCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Авторизация пользователя",
		Run: func(cmd *cobra.Command, args []string) {
			println("Пользователь авторизован!")
		},
	}
}

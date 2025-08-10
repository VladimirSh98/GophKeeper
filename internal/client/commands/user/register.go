package user

import "github.com/spf13/cobra"

// RegisterCmd register user
func RegisterCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "register",
		Short: "Регистрация пользователя",
		Run: func(cmd *cobra.Command, args []string) {
			println("Пользователь зарегистрирован!")
		},
	}
}

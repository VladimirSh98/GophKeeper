package app

import (
	"github.com/VladimirSh98/GophKeeper/internal/client/commands/user"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:  "Gophkeeper",
	Long: `Программа для сохранения различных данных через командную строку`,
}

// Execute init cli
func Execute() {
	rootCmd.AddCommand(user.RegisterCmd())
	rootCmd.AddCommand(user.LoginCmd())
	rootCmd.AddCommand(user.DeleteCmd())
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// go build -o myapp ./cmd/client
// ./myapp --help

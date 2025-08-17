package secrets

import "github.com/spf13/cobra"

// CreateLoginPassCmd create login pass command
func (h *Handler) CreateLoginPassCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "createLoginPassSecret",
		Short: "Добавление логина и пароля",
		Run:   h.secretService.CreateLoginPass,
	}
	cmd.Flags().StringP("login", "l", "", "Логин")
	cmd.Flags().StringP("password", "p", "", "Пароль")
	cmd.Flags().StringArrayP("metadata", "m", []string{}, "Доп инфо")
	cmd.MarkFlagRequired("login")
	cmd.MarkFlagRequired("password")
	return cmd
}

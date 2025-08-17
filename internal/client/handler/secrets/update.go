package secrets

import "github.com/spf13/cobra"

// UpdateCmd update command
func (h *Handler) UpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "updateSecret",
		Short: "Обновление секрета",
		Run:   h.secretService.Update,
	}
	cmd.Flags().IntP("secret", "s", 0, "Номер секрета")
	cmd.Flags().StringP("metadata", "m", "", "Доп инфо")
	cmd.MarkFlagRequired("secret")
	// bank card
	cmd.Flags().StringP("number", "n", "", "Номер карты")
	cmd.Flags().StringP("expiry_month", "exp_m", "", "Месяц срока действия")
	cmd.Flags().StringP("expiry_year", "exp_у", "", "Год срока действия")
	cmd.Flags().StringP("holder_name", "name", "", "Владелец")
	cmd.Flags().StringP("cvv", "cvv", "", "CVV")
	// login pass
	cmd.Flags().StringP("login", "l", "", "Логин")
	cmd.Flags().StringP("password", "pass", "", "Пароль")
	// text
	cmd.Flags().StringP("text", "t", "", "Текст секрета")
	// binary
	cmd.Flags().StringP("binary", "b", "", "Текст секрета")
	return cmd
}

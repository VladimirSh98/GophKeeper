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
	cmd.Flags().StringArrayP("metadata", "m", []string{}, "Доп инфо")
	cmd.MarkFlagRequired("secret")
	// bank card
	cmd.Flags().StringP("number", "n", "", "Номер карты")
	cmd.Flags().StringP("expiry_month", "e", "", "Месяц срока действия")
	cmd.Flags().StringP("expiry_year", "y", "", "Год срока действия")
	cmd.Flags().String("holder_name", "", "Владелец")
	cmd.Flags().StringP("cvv", "c", "", "CVV")
	// login pass
	cmd.Flags().StringP("login", "l", "", "Логин")
	cmd.Flags().StringP("password", "p", "", "Пароль")
	// text
	cmd.Flags().StringP("text", "t", "", "Текст секрета")
	// binary
	cmd.Flags().StringP("binary", "b", "", "Текст секрета")
	return cmd
}

package secrets

import "github.com/spf13/cobra"

// CreateBankCardCmd create bank card command
func (h *Handler) CreateBankCardCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "createBankSecret",
		Short: "Добавление банковской карты",
		Run:   h.secretService.CreateBankCard,
	}
	cmd.Flags().StringP("number", "n", "", "Номер карты")
	cmd.Flags().StringP("expiry_month", "e", "", "Месяц срока действия")
	cmd.Flags().StringP("expiry_year", "y", "", "Год срока действия")
	cmd.Flags().String("holder_name", "", "Владелец")
	cmd.Flags().StringP("cvv", "c", "", "CVV")
	cmd.Flags().StringArrayP("metadata", "m", []string{}, "Доп инфо")
	cmd.MarkFlagRequired("number")
	cmd.MarkFlagRequired("expiry_month")
	cmd.MarkFlagRequired("expiry_year")
	cmd.MarkFlagRequired("holder_name")
	return cmd
}

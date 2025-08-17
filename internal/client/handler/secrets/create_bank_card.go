package secrets

import "github.com/spf13/cobra"

// CreateBankCardCmd create bank card command
func (h *Handler) CreateBankCardCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "createBankSecret",
		Short: "Добавление банковской карты",
		Run:   h.secretService.CreateBinary,
	}
	cmd.Flags().StringP("number", "n", "", "Номер карты")
	cmd.Flags().StringP("expiry_month", "exp_m", "", "Месяц срока действия")
	cmd.Flags().StringP("expiry_year", "exp_у", "", "Год срока действия")
	cmd.Flags().StringP("holder_name", "name", "", "Владелец")
	cmd.Flags().StringP("cvv", "cvv", "", "CVV")
	cmd.Flags().StringP("metadata", "m", "", "Доп инфо")
	cmd.MarkFlagRequired("number")
	cmd.MarkFlagRequired("expiry_month")
	cmd.MarkFlagRequired("expiry_year")
	cmd.MarkFlagRequired("holder_name")
	return cmd
}

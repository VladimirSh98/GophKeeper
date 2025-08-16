package secrets

import "github.com/spf13/cobra"

// CreateBinaryCmd create text command
func (h *Handler) CreateBinaryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "createBinarySecret",
		Short: "Добавление бинарного секрета",
		Run:   h.secretService.CreateBinary,
	}
	cmd.Flags().StringP("binary", "b", "", "Текст секрета")
	cmd.Flags().StringP("metadata", "m", "", "Доп инфо")
	cmd.MarkFlagRequired("binary")
	return cmd
}

package secrets

import "github.com/spf13/cobra"

func (h *Handler) DeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete secret",
		Short: "Удаление секрета",
		Run:   h.secretService.Delete,
	}
	cmd.Flags().StringP("secret", "s", "", "Номер секрета")
	cmd.MarkFlagRequired("secret")
	return cmd
}

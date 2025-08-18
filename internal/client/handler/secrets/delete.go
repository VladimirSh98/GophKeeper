package secrets

import "github.com/spf13/cobra"

func (h *Handler) DeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deleteSecret",
		Short: "Удаление секрета",
		Run:   h.secretService.Delete,
	}
	cmd.Flags().IntP("secret", "s", 0, "Номер секрета")
	cmd.MarkFlagRequired("secret")
	return cmd
}

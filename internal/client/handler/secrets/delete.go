package secrets

import "github.com/spf13/cobra"

func (h *Handler) DeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete user",
		Short: "Удаление пользователя",
		Run:   h.secretService.Delete,
	}
	cmd.Flags().StringP("secret", "s", "", "Номер секрета")
	cmd.MarkFlagRequired("secret")
	return cmd
}

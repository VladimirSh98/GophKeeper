package secrets

import "github.com/spf13/cobra"

func (h *Handler) GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get secrets",
		Short: "Получить все секреты",
		Run:   h.secretService.Get,
	}
	return cmd
}

package secrets

import "github.com/spf13/cobra"

// CreateTextCmd create text command
func (h *Handler) CreateTextCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "createTextSecret",
		Short: "Добавление текстового секрета",
		Run:   h.secretService.CreateText,
	}
	cmd.Flags().StringP("text", "t", "", "Текст секрета")
	cmd.Flags().StringArrayP("metadata", "m", []string{}, "Доп инфо")
	cmd.MarkFlagRequired("text")
	return cmd
}

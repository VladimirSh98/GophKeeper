package user

import "github.com/spf13/cobra"

func (s *Service) Delete(cmd *cobra.Command, args []string) {
	s.logger.Info("Delete called")
}

package user

import "github.com/spf13/cobra"

func (s *Service) Register(cmd *cobra.Command, args []string) {
	s.logger.Info("Register called")
}

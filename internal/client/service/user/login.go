package user

import "github.com/spf13/cobra"

func (s *Service) Login(cmd *cobra.Command, args []string) {
	s.logger.Info("Login called")
}

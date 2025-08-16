package user

import (
	"github.com/VladimirSh98/GophKeeper/internal/client/repository/user"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// Service struct
type Service struct {
	userClient user.ClientInterface
	logger     *zap.Logger
}

// ServiceInterface service interface
type ServiceInterface interface {
	Login(cmd *cobra.Command, args []string)
	Register(cmd *cobra.Command, args []string)
	Delete(cmd *cobra.Command, args []string)
}

// NewService create new service
func NewService(userClient user.ClientInterface, logger *zap.Logger) ServiceInterface {
	return &Service{userClient: userClient, logger: logger}
}

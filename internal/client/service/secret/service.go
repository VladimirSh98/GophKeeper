package secret

import (
	"github.com/VladimirSh98/GophKeeper/internal/client/repository/memory"
	"github.com/VladimirSh98/GophKeeper/internal/client/repository/secret"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// Service struct
type Service struct {
	secretClient secret.ClientInterface
	tokenManager memory.TokenManager
	logger       *zap.Logger
}

// ServiceInterface service interface
type ServiceInterface interface {
	Delete(cmd *cobra.Command, args []string)
	Get(cmd *cobra.Command, args []string)
}

// NewService create new service
func NewService(secretClient secret.ClientInterface, tokenManager memory.TokenManager, logger *zap.Logger) ServiceInterface {
	return &Service{secretClient: secretClient, tokenManager: tokenManager, logger: logger}
}

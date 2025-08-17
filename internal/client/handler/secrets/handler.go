package secrets

import (
	"github.com/VladimirSh98/GophKeeper/internal/client/service/secret"
	"github.com/spf13/cobra"
)

// Handler struct
type Handler struct {
	secretService secret.ServiceInterface
}

// HandlerInterface handler interface
type HandlerInterface interface {
	DeleteCmd() *cobra.Command
	GetCmd() *cobra.Command
	CreateTextCmd() *cobra.Command
	CreateBinaryCmd() *cobra.Command
	CreateBankCardCmd() *cobra.Command
	CreateLoginPassCmd() *cobra.Command
	UpdateCmd() *cobra.Command
}

// NewHandler new user handler
func NewHandler(secretService secret.ServiceInterface) HandlerInterface {
	return &Handler{secretService: secretService}
}

package user

import (
	"github.com/VladimirSh98/GophKeeper/internal/client/service/user"
	"github.com/spf13/cobra"
)

// Handler struct
type Handler struct {
	userService user.ServiceInterface
}

// HandlerInterface handler interface
type HandlerInterface interface {
	LoginCmd() *cobra.Command
	RegisterCmd() *cobra.Command
	DeleteCmd() *cobra.Command
}

// NewHandler new user handler
func NewHandler(userService user.ServiceInterface) HandlerInterface {
	return &Handler{userService: userService}
}

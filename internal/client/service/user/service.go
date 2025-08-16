package user

import "github.com/VladimirSh98/GophKeeper/internal/client/repository/user"

// Service struct
type Service struct {
	userClient user.ClientInterface
}

// ServiceInterface service interface
type ServiceInterface interface{}

// NewService create new service
func NewService(userClient user.ClientInterface) ServiceInterface {
	return &Service{userClient: userClient}
}

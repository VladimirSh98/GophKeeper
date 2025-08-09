package user

import (
	"github.com/VladimirSh98/GophKeeper/internal/server/repository/user"
	authService "github.com/VladimirSh98/GophKeeper/internal/server/service/auth"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Grpc user
type Grpc struct {
	pb.UnimplementedUserServer
	userRepo user.Repository
	logger   *zap.Logger
	auth     authService.Service
}

// NewUserGrpc new grpc user
func NewUserGrpc(userRepo user.Repository, auth authService.Service, logger *zap.Logger) *Grpc {
	return &Grpc{userRepo: userRepo, auth: auth, logger: logger}
}

func (s *Grpc) RegisterService(r grpc.ServiceRegistrar) {
	pb.RegisterUserServer(r, s)
}

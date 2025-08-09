package user

import (
	"github.com/VladimirSh98/GophKeeper/internal/server/repository/user"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"google.golang.org/grpc"
)

// Grpc user
type Grpc struct {
	pb.UnimplementedUserServer
	userRepo user.Repository
}

// NewUserGrpc new grpc user
func NewUserGrpc(userRepo user.Repository) *Grpc {
	return &Grpc{userRepo: userRepo}
}

func (u *Grpc) RegisterService(r grpc.ServiceRegistrar) {
	pb.RegisterUserServer(r, u)
}

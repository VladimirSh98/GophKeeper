package secret

import (
	"github.com/VladimirSh98/GophKeeper/internal/server/repository/secret"
	"github.com/VladimirSh98/GophKeeper/internal/server/repository/user"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Grpc secret
type Grpc struct {
	pb.UnimplementedSecretServer
	userRepo   user.Repository
	secretRepo secret.Repository
	logger     *zap.Logger
}

// NewSecretGrpc new grpc secret
func NewSecretGrpc(userRepo user.Repository, secretRepo secret.Repository, logger *zap.Logger) *Grpc {
	return &Grpc{userRepo: userRepo, secretRepo: secretRepo, logger: logger}
}

// RegisterService register service
func (s *Grpc) RegisterService(r grpc.ServiceRegistrar) {
	pb.RegisterSecretServer(r, s)
}

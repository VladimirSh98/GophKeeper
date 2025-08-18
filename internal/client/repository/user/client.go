package user

import (
	"context"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Client struct
type Client struct {
	client pb.UserClient
	logger *zap.Logger
}

// ClientInterface client interface
type ClientInterface interface {
	Register(ctx context.Context, login string, password string) (*pb.RegisterResponse, error)
	Login(ctx context.Context, login string, password string) (*pb.LoginResponse, error)
	Delete(ctx context.Context, token string) (*pb.DeleteResponse, error)
}

// NewClient create new client
func NewClient(conn *grpc.ClientConn, logger *zap.Logger) ClientInterface {
	client := pb.NewUserClient(conn)
	return &Client{client: client, logger: logger}
}

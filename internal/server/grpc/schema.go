package grpc

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/server/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Server interface
type Server interface {
	Init()
	Start(cancel context.CancelFunc)
	Stop()
	GetServer() *grpc.Server
}

type server struct {
	cfg    *config.Config
	server *grpc.Server
	logger *zap.Logger
}

// NewGrpcServer create grpc server
func NewGrpcServer(logger *zap.Logger, cfg *config.Config) Server {
	newServer := &server{
		logger: logger,
		cfg:    cfg,
	}
	newServer.Init()
	return newServer
}

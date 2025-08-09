package grpc

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/server/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server interface {
	Init()
	Start(cancel context.CancelFunc)
	Stop()
}

type server struct {
	cfg    *config.Config
	server *grpc.Server
	logger *zap.Logger
}

func NewGrpcServer(logger *zap.Logger, cfg *config.Config) Server {
	newServer := &server{
		logger: logger,
		cfg:    cfg,
	}
	newServer.Init()
	return newServer
}

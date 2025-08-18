package grpc

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/server/middleware"
	"google.golang.org/grpc"
	"net"
)

// Start grpc server
func (s *server) Start(cancel context.CancelFunc) {
	lis, err := net.Listen("tcp", s.cfg.ServerAddress)
	if err != nil {
		s.logger.Sugar().Errorf("Server listening failed: %v", err)
	}
	go func() {
		err = s.server.Serve(lis)
		if err != nil {
			s.logger.Sugar().Errorf("Server grpc serve failed: %v", err)
			cancel()
		}
	}()
}

// Stop grpc server
func (s *server) Stop() {
	s.logger.Info("Gracefully stopping gRPC server")
	s.server.GracefulStop()
	s.logger.Info("gRPC server stopped")
}

// Init grpc server
func (s *server) Init() {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.JWTUnaryInterceptor(s.cfg)))
	s.server = grpcServer
}

// GetServer return server
func (s *server) GetServer() *grpc.Server {
	return s.server
}

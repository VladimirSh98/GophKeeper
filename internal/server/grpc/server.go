package grpc

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

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

func (s *server) Stop() {
	s.logger.Info("Gracefully stopping gRPC server")
	s.server.GracefulStop()
	s.logger.Info("gRPC server stopped")
}

func (s *server) Init() {
	grpcServer := grpc.NewServer()
	s.server = grpcServer
}

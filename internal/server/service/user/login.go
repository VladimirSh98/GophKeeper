package user

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/server/middleware"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Grpc) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	userModel, err := s.userRepo.GetUserByLogin(ctx, req.Login, false)
	if err != nil {
		s.logger.Warn("Login GetUserByLogin error", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to get user from db")
	}
	if !middleware.VerifyPassword(req.Password, userModel.Hash) {
		s.logger.Warn("Login GetUserByLogin error", zap.Error(err))
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}
	var token string
	token, err = s.auth.CreateToken(req.Login)
	if err != nil {
		s.logger.Warn("CreateToken error", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to generate token")
	}
	return &pb.LoginResponse{Token: token}, nil
}

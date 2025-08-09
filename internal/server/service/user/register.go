package user

import (
	"context"
	"errors"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Grpc) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hashPass, err := utils.HashPassword(req.GetPassword())
	_, err = s.userRepo.Create(ctx, req.GetLogin(), hashPass)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
		s.logger.Sugar().Warnf("User with login %s already exist\n", req.Login)
		return nil, status.Errorf(codes.FailedPrecondition, "User with login %s already exist", req.Login)
	} else if err != nil {
		s.logger.Sugar().Warnf("User Register error %s", err)
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	var token string
	token, err = s.auth.CreateToken(req.Login)
	if err != nil {
		s.logger.Warn("CreateToken error", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to generate token")
	}
	return &pb.RegisterResponse{Token: token}, nil
}

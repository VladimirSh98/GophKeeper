package user

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Grpc) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	login := ctx.Value(utils.UserLoginKey).(string)
	err := s.userRepo.Delete(ctx, login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed deleting user: %v", err)
	}
	return &pb.DeleteResponse{}, nil
}

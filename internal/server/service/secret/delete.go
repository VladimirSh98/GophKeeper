package secret

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Grpc) Delete(ctx context.Context, req *pb.DeleteSecretRequest) (*pb.SecretModel, error) {
	login := ctx.Value(utils.UserLoginKey).(string)
	_, err := s.secretRepo.DeleteByID(ctx, login, int(req.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, "secret deletion failed")
	}
	return &pb.SecretModel{}, nil
}

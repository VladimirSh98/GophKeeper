package secret

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Grpc) Update(ctx context.Context, req *pb.EditSecretRequest) (*pb.SecretModel, error) {
	login := ctx.Value(utils.UserLoginKey).(string)
	secret, err := s.secretRepo.UpdateByID(ctx, login, int(req.Id), req.Content, req.Metadata)
	if err != nil {
		return nil, status.Error(codes.Internal, "secret update failed")
	}
	return &pb.SecretModel{
		Id:        int64(secret.ID),
		UserId:    int64(secret.UserID),
		Archived:  secret.Archived,
		CreatedAt: timestamppb.New(secret.CreatedAt),
		UpdatedAt: timestamppb.New(secret.UpdatedAt),
		DataType:  pb.DataType(secret.DataType),
		Content:   secret.Content,
		Metadata:  secret.Metadata,
	}, nil
}

package secret

import (
	"context"
	"encoding/json"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Grpc) GetByID(ctx context.Context, req *pb.GetSecretByIDRequest) (*pb.SecretModel, error) {
	login := ctx.Value(utils.UserLoginKey).(string)
	secret, err := s.secretRepo.GetSecretByIDUser(ctx, int(req.Id), login)
	if err != nil {
		return nil, status.Error(codes.Internal, "secret lookup failed")
	}
	var metadata map[string]string
	err = json.Unmarshal(secret.Metadata, &metadata)
	if err != nil {
		return nil, status.Error(codes.Internal, "secret metadata unmarshal failed")
	}
	return &pb.SecretModel{
		Id:        int64(secret.ID),
		UserId:    int64(secret.UserID),
		Archived:  secret.Archived,
		CreatedAt: timestamppb.New(secret.CreatedAt),
		UpdatedAt: timestamppb.New(secret.UpdatedAt),
		DataType:  pb.DataType(secret.DataType),
		Content:   secret.Content,
		Metadata:  metadata,
	}, nil
}

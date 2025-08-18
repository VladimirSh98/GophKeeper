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

func (s *Grpc) Get(ctx context.Context, req *pb.GetSecretsRequest) (*pb.GetSecretsResponse, error) {
	login := ctx.Value(utils.UserLoginKey).(string)
	secrets, err := s.secretRepo.GetSecretsByUser(ctx, login)
	if err != nil {
		return nil, status.Error(codes.Internal, "secret lookup failed")
	}
	response := &pb.GetSecretsResponse{
		Secrets: make([]*pb.SecretModel, 0, len(secrets)),
	}
	for _, secret := range secrets {
		var metadata map[string]string
		err = json.Unmarshal(secret.Metadata, &metadata)
		if err != nil {
			return nil, status.Error(codes.Internal, "secret metadata unmarshal failed")
		}
		response.Secrets = append(response.Secrets, &pb.SecretModel{
			Id:        int64(secret.ID),
			UserId:    int64(secret.UserID),
			Archived:  secret.Archived,
			CreatedAt: timestamppb.New(secret.CreatedAt),
			UpdatedAt: timestamppb.New(secret.UpdatedAt),
			DataType:  pb.DataType(secret.DataType),
			Content:   secret.Content,
			Metadata:  metadata,
		})
	}
	return response, nil
}

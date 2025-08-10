package secret

import (
	"context"
	secretRepo "github.com/VladimirSh98/GophKeeper/internal/server/repository/secret"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Grpc) Create(ctx context.Context, req *pb.CreateSecretRequest) (*pb.SecretModel, error) {
	login := ctx.Value(utils.UserLoginKey).(string)
	user, err := s.userRepo.GetUserByLogin(ctx, login, false)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "user not found")
	}
	var secret secretRepo.Secret
	secret, err = s.secretRepo.Create(ctx, user.ID, int(req.DataType), req.Content, req.Metadata)
	if err != nil {
		return nil, status.Error(codes.Internal, "secret creation failed")
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

package secret

import (
	"context"
	pb "github.com/VladimirSh98/GophKeeper/proto"
)

func (s *Grpc) Delete(ctx context.Context, req *pb.DeleteSecretRequest) (*pb.SecretModel, error) {
	return &pb.SecretModel{}, nil
}

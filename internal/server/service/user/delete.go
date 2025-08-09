package user

import (
	"context"
	pb "github.com/VladimirSh98/GophKeeper/proto"
)

func (s *Grpc) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	print(10)
	return &pb.DeleteResponse{}, nil
}

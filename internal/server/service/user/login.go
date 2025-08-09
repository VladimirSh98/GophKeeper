package user

import (
	"context"
	"fmt"
	pb "github.com/VladimirSh98/GophKeeper/proto"
)

func (u *Grpc) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Print(1000)
	return &pb.LoginResponse{}, nil
}

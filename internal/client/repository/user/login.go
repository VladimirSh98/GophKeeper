package user

import (
	"context"
	pb "github.com/VladimirSh98/GophKeeper/proto"
)

func (c *Client) Login(ctx context.Context, login string, password string) (*pb.LoginResponse, error) {
	response, err := c.client.Login(ctx, &pb.LoginRequest{
		Login:    login,
		Password: password,
	})
	if err != nil {
		c.logger.Sugar().Debugf("Login error: %s", err.Error())
		return nil, err
	}
	return response, nil
}

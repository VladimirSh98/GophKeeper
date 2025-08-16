package user

import (
	"context"
	pb "github.com/VladimirSh98/GophKeeper/proto"
)

// Register method to register user
func (c *Client) Register(ctx context.Context, login string, password string) (*pb.RegisterResponse, error) {
	response, err := c.client.Register(ctx, &pb.RegisterRequest{
		Login:    login,
		Password: password,
	})
	if err != nil {
		c.logger.Sugar().Debugf("Register error: %s", err.Error())
		return nil, err
	}
	return response, nil
}

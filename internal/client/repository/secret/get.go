package secret

import (
	"context"
	pb "github.com/VladimirSh98/GophKeeper/proto"
)

func (c *Client) Get(ctx context.Context) (*pb.GetSecretsResponse, error) {
	response, err := c.client.Get(ctx, &pb.GetSecretsRequest{})
	if err != nil {
		c.logger.Sugar().Debugf("Get secrets error: %s", err.Error())
		return nil, err
	}
	return response, nil
}

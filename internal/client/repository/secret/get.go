package secret

import (
	"context"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"google.golang.org/grpc/metadata"
)

func (c *Client) Get(ctx context.Context, token string) (*pb.GetSecretsResponse, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	response, err := c.client.Get(ctx, &pb.GetSecretsRequest{})
	if err != nil {
		c.logger.Sugar().Debugf("Get secrets error: %s", err.Error())
		return nil, err
	}
	return response, nil
}

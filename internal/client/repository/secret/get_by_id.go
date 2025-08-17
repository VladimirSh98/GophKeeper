package secret

import (
	"context"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"google.golang.org/grpc/metadata"
)

func (c *Client) GetByID(ctx context.Context, token string, secretID int) (*pb.SecretModel, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	response, err := c.client.GetByID(ctx, &pb.GetSecretByIDRequest{Id: int64(secretID)})
	if err != nil {
		c.logger.Sugar().Debugf("Failed to get by id secret model: %s", err.Error())
		return nil, err
	}
	return response, nil
}

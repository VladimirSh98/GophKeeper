package secret

import (
	"context"
	pb "github.com/VladimirSh98/GophKeeper/proto"
)

func (c *Client) Update(
	ctx context.Context,
	secretID int,
	content []byte,
	metadata map[string]string,
) (*pb.SecretModel, error) {
	response, err := c.client.Update(ctx, &pb.EditSecretRequest{
		Id:       int64(secretID),
		Content:  content,
		Metadata: metadata,
	})
	if err != nil {
		c.logger.Sugar().Debugf("Update secret failed with %s", err.Error())
		return nil, err
	}
	return response, nil
}

package secret

import (
	"context"
	pb "github.com/VladimirSh98/GophKeeper/proto"
)

func (c *Client) Delete(ctx context.Context, secretID int) (*pb.SecretModel, error) {
	response, err := c.client.Delete(ctx, &pb.DeleteSecretRequest{Id: int64(secretID)})
	if err != nil {
		c.logger.Sugar().Debugf("Failed to delete secret model: %s", err.Error())
		return nil, err
	}
	return response, nil
}

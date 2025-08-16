package user

import (
	"context"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"google.golang.org/grpc/metadata"
)

func (c *Client) Delete(ctx context.Context, token string) (*pb.DeleteResponse, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	response, err := c.client.Delete(ctx, &pb.DeleteRequest{})
	if err != nil {
		c.logger.Sugar().Debugf("Delete error: %s", err.Error())
		return nil, err
	}
	return response, nil
}

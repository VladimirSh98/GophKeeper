package secret

import (
	"context"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	metadataGrpc "google.golang.org/grpc/metadata"
)

func (c *Client) Create(
	ctx context.Context,
	token string,
	dataType DataType,
	content []byte,
	metadata map[string]string,
) (*pb.SecretModel, error) {
	ctx = metadataGrpc.AppendToOutgoingContext(ctx, "authorization", token)
	response, err := c.client.Create(ctx, &pb.CreateSecretRequest{
		DataType: pb.DataType(dataType),
		Content:  content,
		Metadata: metadata,
	})
	if err != nil {
		c.logger.Sugar().Debugf("Failed to create secret model: %s", err.Error())
		return nil, err
	}
	return response, nil
}

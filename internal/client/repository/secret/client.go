package secret

import (
	"context"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type DataType int

const (
	LoginPassword DataType = iota
	TextData
	BinaryData
	BankCard
)

// Client struct
type Client struct {
	client pb.SecretClient
	logger *zap.Logger
}

// ClientInterface client interface
type ClientInterface interface {
	Delete(ctx context.Context, token string, secretID int) (*pb.SecretModel, error)
	Create(
		ctx context.Context,
		token string,
		dataType DataType,
		content []byte,
		metadata map[string]string,
	) (*pb.SecretModel, error)
	Get(ctx context.Context, token string) (*pb.GetSecretsResponse, error)
	Update(
		ctx context.Context,
		token string,
		secretID int,
		content []byte,
		metadata map[string]string,
	) (*pb.SecretModel, error)
}

// NewClient create new client
func NewClient(conn *grpc.ClientConn, logger *zap.Logger) ClientInterface {
	client := pb.NewSecretClient(conn)
	return &Client{client: client, logger: logger}
}

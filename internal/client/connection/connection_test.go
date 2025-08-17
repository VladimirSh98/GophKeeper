package connection

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/client/config"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

func TestServerConnectionOpenCloseConnection(t *testing.T) {
	listener := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	go func() {
		_ = s.Serve(listener)
	}()
	defer s.Stop()
	cfg := &config.Config{
		ServerAddress: "ecdcv",
	}
	server := &ServerConnection{Cfg: cfg}
	var err error
	server.Conn, err = grpc.DialContext(context.Background(), "ecdcv",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return listener.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	server.CloseConnection()
}

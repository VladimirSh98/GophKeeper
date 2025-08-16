package connection

import (
	"github.com/VladimirSh98/GophKeeper/internal/client/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ServerConnection struct
type ServerConnection struct {
	Conn *grpc.ClientConn
	Cfg  *config.Config
}

// OpenConnection open client connection
func (server *ServerConnection) OpenConnection() error {
	var err error
	server.Conn, err = grpc.Dial(server.Cfg.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	return nil
}

// CloseConnection close client connection
func (server *ServerConnection) CloseConnection() {
	server.Conn.Close()
}

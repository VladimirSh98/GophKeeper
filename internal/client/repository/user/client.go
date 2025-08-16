package user

import "google.golang.org/grpc"

// Client struct
type Client struct {
	Conn *grpc.ClientConn
}

// ClientInterface client interface
type ClientInterface interface{}

// NewClient create new client
func NewClient(conn *grpc.ClientConn) ClientInterface {
	return &Client{Conn: conn}
}

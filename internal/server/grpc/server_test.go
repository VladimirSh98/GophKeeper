package grpc

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/server/config"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestGrpcServer(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	cfg := &config.Config{
		ServerAddress: "localhost:0",
	}
	srv := NewGrpcServer(logger, cfg)
	require.NotNil(t, srv)
	require.NotNil(t, srv.GetServer())
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	go srv.Start(cancel)
	time.Sleep(100 * time.Millisecond)
	srv.Stop()
}

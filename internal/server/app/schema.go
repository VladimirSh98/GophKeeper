package app

import (
	"database/sql"
	grpcServer "github.com/VladimirSh98/GophKeeper/internal/server/grpc"
	"go.uber.org/zap"
)

// App struct
type App struct {
	DB     *sql.DB
	Logger *zap.Logger
	Server grpcServer.Server
}

package app

import (
	"github.com/VladimirSh98/GophKeeper/internal/server/database"
	grpcServer "github.com/VladimirSh98/GophKeeper/internal/server/grpc"
	"go.uber.org/zap"
)

// App struct
type App struct {
	DB     database.DBConnectionStruct
	Logger *zap.Logger
	Server grpcServer.Server
}

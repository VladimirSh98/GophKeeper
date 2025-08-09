package app

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/server/config"
	"github.com/VladimirSh98/GophKeeper/internal/server/database"
	grpcServer "github.com/VladimirSh98/GophKeeper/internal/server/grpc"
	"github.com/VladimirSh98/GophKeeper/internal/server/logger"
	userRepository "github.com/VladimirSh98/GophKeeper/internal/server/repository/user"
	authService "github.com/VladimirSh98/GophKeeper/internal/server/service/auth"
	userGrpc "github.com/VladimirSh98/GophKeeper/internal/server/service/user"
	"log"
)

// NewApp create app
func NewApp(ctx context.Context) (*App, error) {
	initLogger, err := logger.Initialize()
	defer initLogger.Sync()
	if err != nil {
		log.Fatalf("Logger configuration failed: %v", err)
		return nil, err
	}
	cfg := &config.Config{}
	err = config.LoadConfig(cfg)
	if err != nil {
		log.Fatalf("Server configuration failed: %v", err)
		return nil, err
	}
	databaseConn := database.DBConnectionStruct{Cfg: cfg}
	err = databaseConn.OpenConnection()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
		return nil, err
	}
	err = databaseConn.Ping()
	if err != nil {
		log.Fatalf("Database ping failed: %v", err)
		return nil, err
	}
	err = databaseConn.UpgradeMigrations()
	if err != nil {
		log.Printf("Database migrations failed: %v", err)
		return nil, err
	}
	newGrpcServer := grpcServer.NewGrpcServer(initLogger, cfg)

	userRepo := userRepository.NewRepository(databaseConn.Conn)
	auth := authService.NewService(cfg)
	userGrpcService := userGrpc.NewUserGrpc(userRepo, auth, initLogger)
	userGrpcService.RegisterService(newGrpcServer.GetServer())
	return &App{
		Logger: initLogger,
		Server: newGrpcServer,
		DB:     databaseConn,
	}, nil
}

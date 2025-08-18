package app

import (
	"github.com/VladimirSh98/GophKeeper/internal/logger"
	"github.com/VladimirSh98/GophKeeper/internal/server/config"
	"github.com/VladimirSh98/GophKeeper/internal/server/database"
	grpcServer "github.com/VladimirSh98/GophKeeper/internal/server/grpc"
	secretRepository "github.com/VladimirSh98/GophKeeper/internal/server/repository/secret"
	userRepository "github.com/VladimirSh98/GophKeeper/internal/server/repository/user"
	authService "github.com/VladimirSh98/GophKeeper/internal/server/service/auth"
	secretGrpc "github.com/VladimirSh98/GophKeeper/internal/server/service/secret"
	userGrpc "github.com/VladimirSh98/GophKeeper/internal/server/service/user"
	"log"
)

// NewApp create app
func NewApp() (*App, error) {
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
	secretRepo := secretRepository.NewRepository(databaseConn.Conn)
	auth := authService.NewService(cfg)
	userGrpcService := userGrpc.NewUserGrpc(userRepo, secretRepo, auth, initLogger)
	userGrpcService.RegisterService(newGrpcServer.GetServer())
	secretGrpcService := secretGrpc.NewSecretGrpc(userRepo, secretRepo, initLogger)
	secretGrpcService.RegisterService(newGrpcServer.GetServer())

	return &App{
		Logger: initLogger,
		Server: newGrpcServer,
		DB:     databaseConn,
	}, nil
}

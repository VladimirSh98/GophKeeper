package app

import (
	"github.com/VladimirSh98/GophKeeper/internal/client/config"
	"github.com/VladimirSh98/GophKeeper/internal/client/connection"
	secretHandlerRepo "github.com/VladimirSh98/GophKeeper/internal/client/handler/secrets"
	userHandlerRepo "github.com/VladimirSh98/GophKeeper/internal/client/handler/user"
	"github.com/VladimirSh98/GophKeeper/internal/client/repository/memory"
	secretClientRepo "github.com/VladimirSh98/GophKeeper/internal/client/repository/secret"
	userClientRepo "github.com/VladimirSh98/GophKeeper/internal/client/repository/user"
	secretServiceRepo "github.com/VladimirSh98/GophKeeper/internal/client/service/secret"
	userServiceRepo "github.com/VladimirSh98/GophKeeper/internal/client/service/user"
	"github.com/VladimirSh98/GophKeeper/internal/logger"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:  "Gophkeeper",
	Long: `Программа для сохранения различных данных через командную строку`,
}

// Run create cli app
func Run() (App, error) {
	initLogger, err := logger.Initialize()
	defer initLogger.Sync()
	if err != nil {
		log.Fatalf("Logger configuration failed: %v", err)
		return App{}, err
	}
	cfg := &config.Config{}
	err = config.LoadConfig(cfg)
	if err != nil {
		log.Fatalf("Client configuration failed: %v", err)
		return App{}, err
	}
	serverConn := connection.ServerConnection{Cfg: cfg}
	err = serverConn.OpenConnection()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
		return App{}, err
	}
	tokenManager := memory.NewManager(cfg.SecretKey)
	userClient := userClientRepo.NewClient(serverConn.Conn, initLogger)
	secretClient := secretClientRepo.NewClient(serverConn.Conn, initLogger)
	userSevice := userServiceRepo.NewService(userClient, tokenManager, initLogger)
	secretService := secretServiceRepo.NewService(secretClient, tokenManager, initLogger)
	userHandler := userHandlerRepo.NewHandler(userSevice)
	secretHandler := secretHandlerRepo.NewHandler(secretService)
	rootCmd.AddCommand(userHandler.RegisterCmd())
	rootCmd.AddCommand(userHandler.LoginCmd())
	rootCmd.AddCommand(userHandler.DeleteCmd())
	rootCmd.AddCommand(secretHandler.DeleteCmd())
	rootCmd.AddCommand(secretHandler.GetCmd())
	rootCmd.AddCommand(secretHandler.UpdateCmd())
	rootCmd.AddCommand(secretHandler.CreateTextCmd())
	rootCmd.AddCommand(secretHandler.CreateBinaryCmd())
	rootCmd.AddCommand(secretHandler.CreateLoginPassCmd())
	rootCmd.AddCommand(secretHandler.CreateBankCardCmd())
	return App{
		Logger:      initLogger,
		Client:      serverConn,
		RootCommand: rootCmd,
	}, nil
}

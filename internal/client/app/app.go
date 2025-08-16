package app

import (
	"github.com/VladimirSh98/GophKeeper/internal/client/config"
	"github.com/VladimirSh98/GophKeeper/internal/client/connection"
	userHandlerRepo "github.com/VladimirSh98/GophKeeper/internal/client/handler/user"
	userClientRepo "github.com/VladimirSh98/GophKeeper/internal/client/repository/user"
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
	userClient := userClientRepo.NewClient(serverConn.Conn, initLogger)
	userSevice := userServiceRepo.NewService(userClient, initLogger)
	userHandler := userHandlerRepo.NewHandler(userSevice)
	rootCmd.AddCommand(userHandler.RegisterCmd())
	rootCmd.AddCommand(userHandler.LoginCmd())
	rootCmd.AddCommand(userHandler.DeleteCmd())
	return App{}, nil
}

// go build -o myapp ./cmd/client
// ./myapp --help

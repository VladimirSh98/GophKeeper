package main

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/client/app"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()
	application, err := app.Run()
	if err != nil {
		log.Println("Client failed to start: %v", err)
	}
	err = application.RootCommand.ExecuteContext(ctx)
	if err != nil {
		log.Println("Client failed to start: %v", err)
		return
	}
	<-ctx.Done()
	application.Client.CloseConnection()
	application.Logger.Info("Application stopped")
}

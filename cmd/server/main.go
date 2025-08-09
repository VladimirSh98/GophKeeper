package main

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/server/app"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()
	application, err := app.NewApp(ctx)
	if err != nil {
		log.Println("Server failed to start: %v", err)
	}
	application.Server.Start(cancel)
	<-ctx.Done()
	application.Server.Stop()
	application.Logger.Info("Application stopped")
}

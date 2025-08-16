package main

import (
	"github.com/VladimirSh98/GophKeeper/internal/client/app"
	"log"
)

func main() {
	application, err := app.Run()
	if err != nil {
		log.Printf("Client failed to start: %v\n", err)
	}
	defer application.Client.CloseConnection()
	err = application.RootCommand.Execute()
	if err != nil {
		log.Printf("Client failed to command: %v\n", err)
		return
	}
}

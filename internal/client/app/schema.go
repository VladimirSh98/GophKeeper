package app

import (
	"github.com/VladimirSh98/GophKeeper/internal/client/connection"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// App struct
type App struct {
	Client      connection.ServerConnection
	Logger      *zap.Logger
	RootCommand *cobra.Command
}

package logger

import (
	"go.uber.org/zap"
)

// Initialize - initialized logger
func Initialize() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(logger)
	return logger, nil
}

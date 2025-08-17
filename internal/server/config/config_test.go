package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	cfg := &Config{}

	err := LoadConfig(cfg)
	require.NoError(t, err)

	require.Equal(t, "localhost:8000", cfg.ServerAddress)
	require.Equal(t, "postgres://user:pass@localhost:5432/test?sslmode=disable", cfg.DatabaseDSN)
	require.Equal(t, "./migrations", cfg.MigrationsDir)
	require.Equal(t, 55*time.Minute, cfg.TokenExp)
	require.Equal(t, "secret", cfg.SecretKey)
}

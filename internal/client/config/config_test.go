package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg := &Config{}

	err := LoadConfig(cfg)
	require.NoError(t, err)

	require.Equal(t, "localhost:8000", cfg.ServerAddress)
	require.Equal(t, "supersecretkey", cfg.SecretKey)
}

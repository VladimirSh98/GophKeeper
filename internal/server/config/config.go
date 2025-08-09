package config

import (
	"github.com/caarlos0/env/v6"
)

// LoadConfig loads the project configuration
func LoadConfig(cfg *Config) error {
	var err error

	err = env.Parse(cfg)
	if err != nil {
		return err
	}
	return nil
}

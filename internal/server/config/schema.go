package config

// Config struct
type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" json:"server_address" envDefault:"localhost:8000"`
	DatabaseDSN   string `env:"DATABASE_DSN" json:"database_dsn" envDefault:"postgres://user:pass@localhost:5432/test?sslmode=disable"`
	MigrationsDir string `yaml:"migrations_dir" envDefault:"./migrations"`
}

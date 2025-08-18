package config

// Config struct
type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" json:"server_address" envDefault:"localhost:8000"`
	SecretKey     string `env:"CLIENT_SECRET_KEY" envDefault:"supersecretkey"`
}

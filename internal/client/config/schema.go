package config

// Config struct
type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" json:"server_address" envDefault:"localhost:8000"`
	ClientAddress string `env:"CLIENT_ADDRESS" json:"client_address"  envDefault:"localhost:8001"`
}

// config/config.go
package config

type Config struct {
	ServerAddress string
}

func GetConfig() *Config {
	// You can set default values here
	return &Config{
		ServerAddress: ":8080", // Default to listening on port 8080
	}
}

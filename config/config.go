package config

type (
	Config struct {
		Server   ServerConfig
		Database DatabaseConfig
	}

	ServerConfig struct {
		Port int
	}

	DatabaseConfig struct {
		Host       string
		Port       int
		DbName     string
		DbUser     string
		DbPassword string
		SslMode    string
	}
)

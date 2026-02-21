package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	*DbConfig
}

type DbConfig struct {
	DSN string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return &Config{
		&DbConfig{
			DSN: os.Getenv("DSN"),
		},
	}
}

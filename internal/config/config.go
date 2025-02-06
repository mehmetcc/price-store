package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WsUrl string
	Dsn   string
	Port  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		WsUrl: os.Getenv("WS_URL"),
		Dsn:   os.Getenv("DSN"),
		Port:  os.Getenv("PORT"),
	}, nil
}

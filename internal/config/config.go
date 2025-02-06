package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WSUrl string
	Dsn   string
	Port  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		WSUrl: os.Getenv("WS_URL"),
		Dsn:   os.Getenv("DSN"),
		Port:  os.Getenv("PORT"),
	}, nil
}

package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WsUrl    string
	Dsn      string
	Port     string
	ClientId string

	PricerUrl string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		WsUrl:     os.Getenv("WS_URL"),
		Dsn:       os.Getenv("DSN"),
		Port:      os.Getenv("PORT"),
		ClientId:  os.Getenv("CLIENT_ID"),
		PricerUrl: os.Getenv("PRICER_URL"),
	}, nil
}

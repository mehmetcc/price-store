package main

import (
	"log"

	"github.com/mehmetcc/price-store/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	log.Printf("config: %+v", cfg)
}

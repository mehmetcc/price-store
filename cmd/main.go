package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/mehmetcc/price-store/internal/config"
	"github.com/mehmetcc/price-store/internal/db"
	"github.com/mehmetcc/price-store/internal/routes"
	"github.com/mehmetcc/price-store/internal/websocket"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	log.Printf("config: %+v", cfg)

	db.Connect(cfg.Dsn)

	ws, err := websocket.NewClient(cfg.WsUrl)
	if err != nil {
		log.Fatalf("error connecting to websocket: %v", err)
	}
	defer ws.Close()
	go ws.Listen()

	app := fiber.New()
	routes.SetupRoutes(app, ws)
	defer app.Shutdown()

	go func() {
		if err := app.Listen(":" + cfg.Port); err != nil {
			log.Fatalf("Error starting Fiber server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down goodby u piece of shit")
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehmetcc/price-store/internal/websocket"
)

func SetupRoutes(app *fiber.App, wsClient *websocket.Client) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})
	app.Post("/symbol", func(c *fiber.Ctx) error {
		type request struct {
			Symbol string `json:"symbol"`
		}

		var req request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse request body",
			})
		}
		if req.Symbol == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "symbol is required",
			})
		}

		if err := wsClient.SendSymbol(req.Symbol); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to send symbol to websocket",
			})
		}

		return c.JSON(fiber.Map{
			"status": "tracking initiated",
			"symbol": req.Symbol,
		})
	})
}

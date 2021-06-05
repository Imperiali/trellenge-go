package main

import (
	"fmt"

	"github.com/Pelegrinetti/trellenge-go/internal/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.New()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Listen(fmt.Sprintf(":%d", config.Port))
}

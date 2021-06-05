package http

import (
	"fmt"

	"github.com/Pelegrinetti/trellenge-go/pkg/helloworld"
	"github.com/gofiber/fiber/v2"
)

type Server interface {
	Run(port int)
}

type server struct {
	port   int
	config fiber.Config
}

func (s server) Run(port int) {
	app := fiber.New(s.config)

	app.Get("/", func(c *fiber.Ctx) error {
		msg := helloworld.GetMessage()

		return c.SendString(msg)
	})

	app.Listen(fmt.Sprintf(":%d", port))
}

func New(config fiber.Config) Server {
	return server{
		config: config,
	}
}

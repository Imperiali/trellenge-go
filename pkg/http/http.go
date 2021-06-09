package http

import (
	"fmt"
	"github.com/Pelegrinetti/trellenge-go/pkg/container"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server interface {
	Run(port int)
}

type server struct {
	port      int
	container *container.Container
}

func (s *server) Run(port int) {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/users", GetUser(s.container))

	app.Post("/users", CreateUser(s.container))

	app.Get("/games", GetGame(s.container))

	app.Post("/games", CreateGame(s.container))

	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}

func New(ctn *container.Container) Server {
	return &server{
		container: ctn,
	}
}

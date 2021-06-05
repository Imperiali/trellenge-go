package main

import (
	"github.com/Pelegrinetti/trellenge-go/internal/config"
	"github.com/Pelegrinetti/trellenge-go/pkg/http"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.New()
	s := http.New(fiber.Config{})

	s.Run(config.Port)
}

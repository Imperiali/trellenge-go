package http

import (
	"fmt"
	"github.com/Pelegrinetti/trellenge-go/pkg/container"
	"github.com/Pelegrinetti/trellenge-go/pkg/games"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func GetGame(ctn *container.Container) fiber.Handler {
	return func(c *fiber.Ctx) error {
		title := c.Query("title")

		game := games.Game{
			Title: title,
		}

		data, cacheUserError := game.GetFromCache(ctn.Cache)
		if cacheUserError != redis.Nil && cacheUserError != nil {
			fmt.Println("Error getting game.", cacheUserError.Error())
			return c.SendStatus(500)
		}

		if err := game.Unmarshal(data); err != nil {
			fmt.Println("Error decoding game.", err.Error(), data, title)
			return c.SendStatus(500)
		}

		return c.Status(200).JSON(game)
	}
}

func CreateGame(ctn *container.Container) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var game games.Game

		if err := game.Unmarshal(c.Body()); err != nil {
			fmt.Println("Error parsing game.", err.Error())
			return c.SendStatus(400)
		}

		data, cacheUserError := game.GetFromCache(ctn.Cache)
		if cacheUserError != redis.Nil && cacheUserError != nil {
			fmt.Println("Error getting game.", cacheUserError.Error())
			return c.SendStatus(500)
		}

		if data != redis.Nil && data != "" {
			return c.SendStatus(201)
		}

		if err := game.Create(ctn.Cache); err != nil {
			fmt.Println("Error creating game in cache.", err.Error())
			return c.SendStatus(500)
		}

		return c.Status(201).JSON(game)
	}
}

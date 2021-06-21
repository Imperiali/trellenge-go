package http

import (
	"fmt"
	"strconv"

	"github.com/Pelegrinetti/trellenge-go/pkg/container"
	"github.com/Pelegrinetti/trellenge-go/pkg/games"
	"github.com/Pelegrinetti/trellenge-go/pkg/users"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func GetUser(ctn *container.Container) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Query("name")
		age, strConvAgeError := strconv.Atoi(c.Query("age"))
		if strConvAgeError != nil {
			fmt.Println("Error parsing user age.", strConvAgeError.Error())
			return c.SendStatus(400)
		}

		user := users.User{
			Name: name,
			Age:  age,
		}

		data, cacheUserError := user.GetUserFromCache(ctn.Cache)
		if cacheUserError != redis.Nil && cacheUserError != nil {
			fmt.Println("Error getting user.", cacheUserError.Error())
			return c.SendStatus(500)
		}

		if err := user.Unmarshal(data); err != nil {
			fmt.Println("Error decoding user.", err.Error(), data, name, age)
			return c.SendStatus(500)
		}

		return c.Status(200).JSON(user)
	}
}

func CreateUser(ctn *container.Container) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user users.User

		if err := user.Unmarshal(c.Body()); err != nil {
			fmt.Println("Error parsing user.", err.Error())
			return c.SendStatus(400)
		}

		data, cacheUserError := user.GetUserFromCache(ctn.Cache)
		if cacheUserError != redis.Nil && cacheUserError != nil {
			fmt.Println("Error getting user.", cacheUserError.Error())
			return c.SendStatus(500)
		}

		for i, item := range user.Games {
			g := games.New(item.Title, item.Category, []string{})
			gameData, gameDataErr := g.GetFromCache(ctn.Cache)

			if gameData == "" || gameData == nil || gameData == redis.Nil {
				fmt.Println("Game doesn't exists")
				return c.SendStatus(404)
			}

			if gameDataErr != nil && gameDataErr != redis.Nil {
				fmt.Println("Error getting from cache: ", gameDataErr)
				return gameDataErr
			}

			if err := g.Unmarshal(gameData); err != nil {
				fmt.Println("Error unmarshing data: ", err)
				return c.SendStatus(500)
			}

			if !g.ContainsUserId(user.GetCacheKey()) {
				g.UserIds = append(g.UserIds, user.GetCacheKey())
			}

			user.Games[i] = g

			if err := g.Create(ctn.Cache); err != nil {
				fmt.Println("Error updating game: ", err)
				return c.SendStatus(500)
			}
		}

		if data != redis.Nil && data != "" {
			return c.SendStatus(201)
		}

		if _, err := user.SetUserInCache(ctn.Cache); err != nil {
			fmt.Println("Error creating user in cache.", err.Error())
			return c.SendStatus(500)
		}

		return c.Status(201).JSON(user)
	}
}

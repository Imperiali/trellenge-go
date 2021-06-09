package http

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Pelegrinetti/trellenge-go/pkg/container"
	"github.com/Pelegrinetti/trellenge-go/pkg/users"
	"github.com/go-redis/redis/v8"
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

	app.Get("/users", func(c *fiber.Ctx) error {
		var user users.User

		name := c.Query("name")
		age, strConvAgeError := strconv.Atoi(c.Query("age"))
		if strConvAgeError != nil {
			fmt.Println("Error parsing user age.", strConvAgeError.Error())
			return c.SendStatus(400)
		}

		cacheKey := fmt.Sprintf("%s::%d", name, age)

		data, cacheUserError := s.container.Cache.Get(cacheKey)
		if cacheUserError != redis.Nil && cacheUserError != nil {
			fmt.Println("Error getting user.", cacheUserError.Error())
			return c.SendStatus(500)
		}

		if err := json.Unmarshal([]byte(data.(string)), &user); err != nil {
			fmt.Println("Error decoding user.", err.Error(), data, name, age)
			return c.SendStatus(500)
		}

		return c.Status(200).JSON(user)
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var user users.User

		if err := json.Unmarshal(c.Body(), &user); err != nil {
			fmt.Println("Error parsing user.", err.Error())
			return c.SendStatus(400)
		}

		cacheKey := fmt.Sprintf("%s::%d", user.Name, user.Age)

		data, cacheUserError := s.container.Cache.Get(cacheKey)
		if cacheUserError != redis.Nil && cacheUserError != nil {
			fmt.Println("Error getting user.", cacheUserError.Error())
			return c.SendStatus(500)
		}

		if data != redis.Nil && data != "" {
			return c.SendStatus(201)
		}

		encodedUser, encodedUserError := json.Marshal(user)
		if encodedUserError != nil {
			fmt.Println("Error encoding user.", encodedUserError.Error())
			return c.SendStatus(500)
		}

		if _, err := s.container.Cache.Set(cacheKey, encodedUser, time.Minute*5); err != nil {
			fmt.Println("Error creating user in cache.", err.Error())
			return c.SendStatus(500)
		}

		return c.Status(201).JSON(user)
	})

	app.Listen(fmt.Sprintf(":%d", port))
}

func New(ctn *container.Container) Server {
	return &server{
		container: ctn,
	}
}

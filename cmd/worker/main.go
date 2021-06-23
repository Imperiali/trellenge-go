package main

import (
	"fmt"

	"github.com/Pelegrinetti/trellenge-go/internal/config"
	"github.com/Pelegrinetti/trellenge-go/pkg/cache"
	"github.com/Pelegrinetti/trellenge-go/pkg/container"
	"github.com/Pelegrinetti/trellenge-go/pkg/games"
	"github.com/Pelegrinetti/trellenge-go/pkg/http"
	"github.com/Pelegrinetti/trellenge-go/pkg/messenger"
	"github.com/Pelegrinetti/trellenge-go/pkg/users"
	"github.com/nsqio/go-nsq"
)

func handleWithDeletedGames(c cache.Cache) nsq.HandlerFunc {
	return func(m *nsq.Message) error {
		if len(m.Body) == 0 {
			return nil
		}

		var game games.Game

		if err := game.Unmarshal(m.Body); err != nil {
			return err
		}

		for _, userId := range game.UserIds {
			var user users.User

			fmt.Println(userId)

			data, _ := c.Get(userId)
			user.Unmarshal(data)
			user.RemoveGame(game.Title)
			user.SetUserInCache(c)

			fmt.Println("Deleted games")
		}

		return nil
	}
}

func ListenDeletedGames(c cache.Cache) (error, func()) {
	msgr := messenger.New()

	consumer, consumerError := msgr.CreateConsumer("deletedGames", "games")

	if consumerError != nil {
		fmt.Println("Cannot start worker:", consumerError.Error())

		return consumerError, func() { consumer.Stop() }
	}

	consumer.AddConcurrentHandlers(handleWithDeletedGames(c), 1)

	err := consumer.ConnectToNSQLookupd("nsq.hud:4161")

	if err != nil {
		fmt.Println(err)
		return err, func() { consumer.Stop() }

	}

	return nil, func() { consumer.Stop() }

}

func main() {
	config := config.New()
	cacheClient := cache.New(config.CacheAddress, config.CachePassword)
	ctn := container.New()

	ctn.WithCache(cacheClient)
	_, stopConsumer := ListenDeletedGames(ctn.Cache)

	defer stopConsumer()

	s := http.New(ctn)

	s.Run(3334)
}

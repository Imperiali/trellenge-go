package games

import (
	"encoding/json"
	"fmt"
	"github.com/Pelegrinetti/trellenge-go/pkg/cache"
	"reflect"
	"time"
)

type Game struct {
	Title    string   `json:"title"`
	Category string   `json:"category"`
	UserId   []string `json:"user_id"`
}

func (g *Game) Unmarshal(data interface{}) error {
	if reflect.TypeOf(data).String() == "[]uint8" {
		return json.Unmarshal(data.([]byte), g)
	}

	return json.Unmarshal([]byte(data.(string)), g)
}

func (g *Game) getCacheKey() string {
	return g.Title
}

func (g *Game) GetFromCache(c cache.Cache) (interface{}, error) {
	cacheKey := g.getCacheKey()

	return c.Get(cacheKey)
}

func (g *Game) Create(c cache.Cache) error {
	cacheKey := g.getCacheKey()

	encodedGame, encodedGameError := json.Marshal(g)
	if encodedGameError != nil {
		fmt.Println("Error encoding user.", encodedGameError.Error())
		return encodedGameError
	}
	_, err := c.Set(cacheKey, encodedGame, time.Minute*5)

	return err
}

func (g *Game) ContainsUserId(id string) bool {
	for _, userId := range g.UserId {
		if userId == id {
			return true
		}
	}
	return false
}

func New(title, category string, userId []string) Game {
	return Game{
		Title:    title,
		Category: category,
		UserId:   userId,
	}
}

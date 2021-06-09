package users

import (
	"encoding/json"
	"fmt"
	"github.com/Pelegrinetti/trellenge-go/pkg/cache"
	"github.com/Pelegrinetti/trellenge-go/pkg/games"
	"reflect"
	"time"
)

type User struct {
	Name  string       `json:"name"`
	Age   int          `json:"age"`
	Games []games.Game `json:"games"`
}

func (u *User) Unmarshal(data interface{}) error {
	if reflect.TypeOf(data).String() == "[]uint8" {
		return json.Unmarshal(data.([]byte), u)
	}

	return json.Unmarshal([]byte(data.(string)), u)
}

func (u *User) GetCacheKey() string {
	return fmt.Sprintf("%s::%d", u.Name, u.Age)
}

func (u *User) GetUserFromCache(c cache.Cache) (interface{}, error) {
	cacheKey := u.GetCacheKey()

	return c.Get(cacheKey)
}

func (u *User) SetUserInCache(c cache.Cache) (string, error) {
	cacheKey := u.GetCacheKey()

	encodedUser, encodedUserError := json.Marshal(u)
	if encodedUserError != nil {
		fmt.Println("Error encoding user.", encodedUserError.Error())
		return "Error", encodedUserError
	}

	return c.Set(cacheKey, encodedUser, time.Minute*5)
}

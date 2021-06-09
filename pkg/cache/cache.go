package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, data interface{}, ttl time.Duration) (string, error)
}

type cache struct {
	client *redis.Client
}

func (c *cache) Get(key string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, err := c.client.Get(ctx, key).Result()

	return result, err
}

func (c *cache) Set(key string, data interface{}, ttl time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, err := c.client.Set(ctx, key, data, ttl).Result()

	return result, err
}

func New(address, password string) Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	return &cache{
		client,
	}
}

package main

import (
	"github.com/Pelegrinetti/trellenge-go/internal/config"
	"github.com/Pelegrinetti/trellenge-go/pkg/cache"
	"github.com/Pelegrinetti/trellenge-go/pkg/container"
	"github.com/Pelegrinetti/trellenge-go/pkg/http"
)

func main() {
	config := config.New()
	cacheClient := cache.New(config.CacheAddress, config.CachePassword)
	ctn := container.New()

	ctn.WithCache(cacheClient)

	s := http.New(ctn)

	s.Run(config.Port)
}

package container

import (
	"github.com/Pelegrinetti/trellenge-go/pkg/cache"
	"github.com/Pelegrinetti/trellenge-go/pkg/db"
)

type Container struct {
	Cache    cache.Cache
	Database db.Database
}

func (c *Container) WithCache(cacheClient cache.Cache) {
	c.Cache = cacheClient
}

func New() *Container {
	database := db.New()
	return &Container{
		Database: database,
	}
}

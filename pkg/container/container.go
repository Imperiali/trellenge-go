package container

import "github.com/Pelegrinetti/trellenge-go/pkg/cache"

type Container struct {
	Cache cache.Cache
}

func (c *Container) WithCache(cacheClient cache.Cache) {
	c.Cache = cacheClient
}

func New() *Container {
	return &Container{}
}

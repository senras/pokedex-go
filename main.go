package main

import (
	"time"

	"github.com/senras/pokedex-go/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	cfg := &config{
		cache: cache,
	}
	initConfig(cfg)
	startRepl(cfg)
}

package main

import (
	"github.com/elifortmort/pokedex/internal/pokecache"
	"time"
)

func main() {
	t := 5 * time.Second
	cache := pokecache.NewCache(t)

	c := config{
		commands: getCommands(),
		cache: cache,
		pokedex: make(map[string]Pokemon),
	}
	replLoop(&c)
}
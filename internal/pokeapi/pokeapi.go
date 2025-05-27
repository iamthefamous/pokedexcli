package pokeapi

import (
	"github.com/iamthefamous/pokedexcli.git/internal/pokecache"
	"time"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

var apiCache = pokecache.NewCache(5 * time.Second)

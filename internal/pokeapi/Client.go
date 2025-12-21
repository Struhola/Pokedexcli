package pokeapi

import (
	"Pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration, cache pokecache.Cache) Client {
	return Client{
		cache: cache,
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

package pokeapi

import (
	"go-http-client/internal/pokecache"
	"net/http"
	"time"
)

// Client -
type Client struct {
	cache     pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}

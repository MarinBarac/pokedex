package pokeapi

import (
	"io"
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(5*time.Minute, nil),
	}
}

func (client *Client) GETRequest(url string) ([]byte, error) {
	existingData, ok := client.cache.Get(url)

	if ok {
		return existingData, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return []byte{}, err
	}

	resp, err := client.httpClient.Do(req)

	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	byteData, err := io.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	client.cache.Add(url, byteData)

	return byteData, nil
}

package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListPokemons(name string) (RespShallowPokemons, error) {
	url := baseURL + "/location-area/" + name

	data, err := c.GETRequest(url)

	if err != nil {
		return RespShallowPokemons{}, err
	}

	var result RespShallowPokemons

	if err := json.Unmarshal(data, &result); err != nil {
		return RespShallowPokemons{}, err
	}

	return result, nil
}

package pokeapi

import "encoding/json"

func (c *Client) PokemonDetails(name string) (RespShallowPokemon, error) {
	url := baseURL + "/pokemon/" + name
	data, err := c.GETRequest(url)

	if err != nil {
		return RespShallowPokemon{}, err
	}

	var pokemon RespShallowPokemon

	if err = json.Unmarshal(data, &pokemon); err != nil {
		return RespShallowPokemon{}, err
	}

	return pokemon, nil
}

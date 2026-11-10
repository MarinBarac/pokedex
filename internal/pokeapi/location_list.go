package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListLocations(pageUrl *string) (ResShallowLocations, error) {
	url := baseURL + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	respData, err := c.GETRequest(url)

	if err != nil {
		return ResShallowLocations{}, err
	}

	shallowLocations := ResShallowLocations{}

	if err := json.Unmarshal(respData, &shallowLocations); err != nil {
		return ResShallowLocations{}, err
	}

	return shallowLocations, nil
}

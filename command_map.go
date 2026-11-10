package main

import (
	"fmt"
)

type mapState struct {
	NextURL     string
	PreviousURL string
}

var mState = mapState{
	NextURL:     "https://pokeapi.co/api/v2/location-area/",
	PreviousURL: "",
}

func commandMap(conf *config, args ...string) error {
	response, err := conf.pokeapiClient.ListLocations(conf.nextLocationsURL)

	if err != nil {
		return err
	}

	conf.prevLocationsURL = response.Previous
	conf.nextLocationsURL = response.Next

	for _, loc := range response.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(conf *config, args ...string) error {
	if conf.prevLocationsURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	response, err := conf.pokeapiClient.ListLocations(conf.prevLocationsURL)

	if err != nil {
		return err
	}

	for _, loc := range response.Results {
		fmt.Println(loc.Name)
	}

	conf.nextLocationsURL = response.Next
	conf.prevLocationsURL = response.Previous

	return nil
}

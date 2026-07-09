package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationAreas, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreas{}, fmt.Errorf(
			"PokeAPI returned a non-OK status code: %d", res.StatusCode)
	}

	var locAreas LocationAreas
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locAreas); err != nil {
		return LocationAreas{}, err
	}

	return locAreas, nil
}

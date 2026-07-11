package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"git.hwanimation.tech/henrikwilhelmsen/pokedex/internal/pokecache"
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

func GetLocationAreas(url string, cache *pokecache.Cache) (LocationAreas, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	}

	var data []byte

	if cached, ok := cache.Get(url); ok {
		data = cached
	} else {
		res, err := http.Get(url)
		if err != nil {
			return LocationAreas{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return LocationAreas{}, fmt.Errorf(
				"PokeAPI returned a non-OK status code: %d", res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationAreas{}, err
		}
		cache.Add(url, body)
		data = body
	}

	var locAreas LocationAreas
	if err := json.Unmarshal(data, &locAreas); err != nil {
		return LocationAreas{}, err
	}

	return locAreas, nil
}

package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"git.hwanimation.tech/henrikwilhelmsen/pokedex/internal/pokecache"
)

// LocationAreas represents a list of location areas from the PokeAPI.
type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// LocationAreaDetails represents the details of a location area from the PokeAPI.
type LocationAreaDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

// GetLocationAreas retrieves a list of location areas from the PokeAPI.
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

// GetLocationAreaDetails retrieves details of a location area from the PokeAPI.
func GetLocationAreaDetails(name string, cache *pokecache.Cache) (LocationAreaDetails, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)

	var data []byte

	if cached, ok := cache.Get(url); ok {
		data = cached
	} else {
		res, err := http.Get(url)
		if err != nil {
			return LocationAreaDetails{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return LocationAreaDetails{}, fmt.Errorf(
				"PokeAPI returned a non-OK status code: %d", res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaDetails{}, err
		}
		cache.Add(url, body)
		data = body
	}

	var locAreaDetails LocationAreaDetails
	if err := json.Unmarshal(data, &locAreaDetails); err != nil {
		return LocationAreaDetails{}, err
	}

	return locAreaDetails, nil
}

// GetLocationAreaPokemon retrieves the list of Pokemon that can be encountered in a location area.
func GetLocationAreaPokemon(location_name string, cache *pokecache.Cache) ([]string, error) {
	areaDetails, err := GetLocationAreaDetails(location_name, cache)
	if err != nil {
		return nil, fmt.Errorf("Failed to get location details: %v", err)
	}

	var pokemonNames []string
	for _, encounter := range areaDetails.PokemonEncounters {
		pokemonNames = append(pokemonNames, encounter.Pokemon.Name)
	}
	return pokemonNames, nil
}

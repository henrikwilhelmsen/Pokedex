package cli

import (
	"fmt"
	"math/rand"
	"os"
	"sort"

	"git.hwanimation.tech/henrikwilhelmsen/pokedex/internal/pokeapi"
	"git.hwanimation.tech/henrikwilhelmsen/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	nextUrl     string
	previousUrl string
	cache       *pokecache.Cache
	pokedex     map[string]pokeapi.Pokemon
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught Pokemon",
			callback:    CommandInspect,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch the given Pokemon",
			callback:    CommandCatch,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call displays the next 20 locations.",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world.",
			callback:    CommandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location and list the possible Pokemon encounters.",
			callback:    CommandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    CommandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
	}

}

// CommandInspect prints detailed information for Pokemon in the Pokedex
func CommandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: inspect <caught_pokemon>")
	}

	pokemonName := args[0]
	pokemonDetails, ok := cfg.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("You have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemonDetails.Name)
	fmt.Printf("Height: %d\n", pokemonDetails.Height)
	fmt.Printf("Weight: %d\n", pokemonDetails.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemonDetails.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokeType := range pokemonDetails.Types {
		fmt.Printf("  -%s\n", pokeType.Type.Name)
	}

	return nil
}

// CommandCatch tries catching the given Pokemon and reports the result. If successful,
// the Pokemon is added to the Pokedex.
func CommandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: catch <pokemon>")
	}

	pokemonName := args[0]
	pokemonDetails, err := pokeapi.GetPokemon(pokemonName, cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	caught := rand.Intn(pokemonDetails.BaseExperience/20) == 1
	if caught {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemonDetails
	} else {
		fmt.Printf("%s broke free!\n", pokemonName)
	}

	return nil
}

// CommandExplore explores a location and lists the possible Pokemon encounters.
func CommandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: explore <location>")
	}
	location := args[0]
	pokemon, err := pokeapi.GetLocationAreaPokemon(location, cfg.cache)
	if err != nil {
		return err
	}

	for _, p := range pokemon {
		fmt.Println(p)
	}
	return nil
}

// CommandMap displays the next 20 location areas in the Pokemon world.
func CommandMap(cfg *config, _ ...string) error {
	locAreas, err := pokeapi.GetLocationAreas(cfg.nextUrl, cfg.cache)
	if err != nil {
		return err
	}

	cfg.previousUrl = locAreas.Previous
	cfg.nextUrl = locAreas.Next

	for _, loc := range locAreas.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// CommandMapBack displays the previous 20 location areas in the Pokemon world.
func CommandMapBack(cfg *config, _ ...string) error {
	locAreas, err := pokeapi.GetLocationAreas(cfg.previousUrl, cfg.cache)
	if err != nil {
		return err
	}

	cfg.previousUrl = locAreas.Previous
	cfg.nextUrl = locAreas.Next

	for _, loc := range locAreas.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// CommandExit exits the Pokedex application.
func CommandExit(cfg *config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// CommandHelp displays the help message for the Pokedex application.
func CommandHelp(cfg *config, _ ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	// Create a sorted slice with the dict keys so we can print the commands in order
	cmds := GetCommands()
	cmdKeys := []string{}
	for k := range cmds {
		cmdKeys = append(cmdKeys, k)
	}
	sort.Strings(cmdKeys)

	// Print name and description for all of the commands
	for _, k := range cmdKeys {
		cmd := cmds[k] // Access by key using the sorted slice for ordering
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

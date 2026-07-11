package cli

import (
	"fmt"
	"os"
	"sort"

	"git.hwanimation.tech/henrikwilhelmsen/pokedex/internal/pokeapi"
	"git.hwanimation.tech/henrikwilhelmsen/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	nextUrl     string
	previousUrl string
	cache       *pokecache.Cache
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
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

func CommandMap(cfg *config) error {
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

func CommandMapBack(cfg *config) error {
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

func CommandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(cfg *config) error {
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

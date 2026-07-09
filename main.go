package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

// Split string on whitespace, removing empty strings and converting to lowercase
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmd, ok := getCommands()[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
	}
}

func main() {
	startRepl()
}

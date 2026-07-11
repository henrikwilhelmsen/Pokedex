package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"git.hwanimation.tech/henrikwilhelmsen/pokedex/internal/pokecache"
)

// Split string on whitespace, removing empty strings and converting to lowercase
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

// Start the Pokedex REPL loop
func StartRepl() {
	const duration = 60 * time.Second
	cache := pokecache.NewCache(duration)
	cfg := &config{cache: cache}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmd, ok := GetCommands()[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(cfg); err != nil {
			fmt.Println(err)
			continue
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}
	}
}

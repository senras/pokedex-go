package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/senras/pokedex-go/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

type config struct {
	cliCommands map[string]cliCommand
	nextUrl     string
	previousUrl string
	cache       pokecache.Cache
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	for i, word := range words {
		words[i] = strings.Trim(word, ".,!?;:")
	}
	return words
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pokedex > ")
		if !reader.Scan() {
			if err := reader.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}
			return
		}

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		if command, exists := cfg.cliCommands[words[0]]; exists {
			err := command.callback(cfg, words[1:]...)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
		fmt.Println()

	}
}

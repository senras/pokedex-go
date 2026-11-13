package main

import "github.com/senras/pokedex-go/internal/pokeapi"

func initConfig(cfg *config) {
	cfg.cliCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    func(cfg *config, args ...string) error { return commandExit() },
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func(cfg *config, args ...string) error { return commandHelp(cfg) },
		},
		"map": {
			name:        "map",
			description: "Displays a map of the Pokemon world",
			callback:    func(cfg *config, args ...string) error { return commandMap(cfg) },
		},
		"mapb": {
			name:        "mapb",
			description: "Displays a map of the Pokemon world (previous if available)",
			callback:    func(cfg *config, args ...string) error { return commandMapb(cfg) },
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of all Pokemon found in a given location",
			callback:    func(cfg *config, args ...string) error { return commandExplore(cfg, args...) },
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a Pokemon by name",
			callback:    func(cfg *config, args ...string) error { return commandCatch(cfg, args...) },
		},
	}
	cfg.pokemonsCatched = make(map[string]pokeapi.Pokemon)
}

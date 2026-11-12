package main

import (
	"fmt"

	"github.com/senras/pokedex-go/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	var location string
	if len(args) > 0 {
		location = args[0]
	} else {
		fmt.Println("No location entered. Please try again.")
		return nil
	}
	pokemonList, err := pokeapi.ListPokemonByLocation(location, &cfg.cache)

	if err != nil {
		fmt.Printf("Error fetching Pokemons for %s: %v\n", location, err)
		return nil
	}
	if len(pokemonList) == 0 {
		fmt.Println("No pokemon found for:", location)
		return nil
	}
	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonList {
		fmt.Println("- ", pokemon)
	}

	return nil
}

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/senras/pokedex-go/internal/pokeapi"
)

func commandExplore(cfg *config) error {

	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a location to explore (e.g., 'kanto', 'johto', 'hoenn', etc.):")
	if !reader.Scan() {
		if err := reader.Err(); err != nil {
			fmt.Println("Error reading input:", err)
		}
		return nil
	}
	location := reader.Text()
	if len(location) == 0 {
		fmt.Println("No location entered. Please try again.")
		return nil
	}
	pokemonList, err := pokeapi.ListPokemonByLocation(&location, &cfg.cache)

	if err != nil {
		fmt.Println("Error fetching Pokemons for %w: %w", location, err)
		return nil
	}
	if len(pokemonList) == 0 {
		fmt.Println("No pokemon found for:", location)
		return nil
	}
	fmt.Println("Pokemon found in", location, ":")
	for _, pokemon := range pokemonList {
		fmt.Println("*", pokemon)
	}

	return nil
}

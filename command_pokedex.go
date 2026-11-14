package main

import (
	"fmt"
)

func commandPokedex(cfg *config) error {

	fmt.Printf("Your Pokedex: \n")
	if len(cfg.pokemonsCatched) == 0 {
		fmt.Printf("No pokemons catched yet")
	}
	for pokemonName := range cfg.pokemonsCatched {
		fmt.Printf("- %s\n", pokemonName)
	}
	return nil
}

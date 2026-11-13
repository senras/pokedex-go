package main

import (
	"fmt"
	"math/rand"

	"github.com/senras/pokedex-go/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	var pokemonName string
	if len(args) == 1 {
		pokemonName = args[0]
	} else {
		fmt.Println("No Pokemon name to catch entered. Please try again.")
		return nil
	}

	pokemon, err := pokeapi.GetPokemonByName(pokemonName)
	if err != nil {
		fmt.Printf("Error fetching information of Pokemon %s", pokemonName)
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	throwResult := throwPokeball(pokemon)

	if throwResult {
		fmt.Printf("%s was caught!!\n", pokemon.Name)
		cfg.pokemonsCatched[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil

}

func throwPokeball(pokemon pokeapi.Pokemon) bool {

	probability := 100.0 / (100.0 + float64(pokemon.BaseExperience))
	return rand.Float64() < probability
}

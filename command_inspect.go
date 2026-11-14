package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	var pokemonName string
	if len(args) == 1 {
		pokemonName = args[0]
	} else {
		fmt.Println("The Pokemon name entered doesn't exist. Please try again.")
		return nil
	}

	pokemon, ok := cfg.pokemonsCatched[pokemonName]
	if ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats: \n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types: ")
		for _, t := range pokemon.Types {
			fmt.Printf("%s ", t.Type.Name)
		}
		fmt.Println()
	} else {
		fmt.Printf("you have not caught that pokemon")
	}

	return nil

}

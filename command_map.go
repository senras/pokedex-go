package main

import (
	"fmt"

	"github.com/senras/pokedex-go/internal/pokeapi"
)

const (
	locationAreaURL = "https://pokeapi.co/api/v2/location-area"
)

func commandMap(cfg *config) error {
	url := locationAreaURL
	if cfg.nextUrl != "" {
		url = cfg.nextUrl
	}

	locationsResp, err := pokeapi.ListLocationAreas(&url)
	if err != nil {
		return err
	}

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	cfg.nextUrl = locationsResp.Next
	cfg.previousUrl = locationsResp.Previous
	return nil
}

func commandMapb(cfg *config) error {
	url := locationAreaURL
	if cfg.previousUrl != "" {
		url = cfg.previousUrl
	}

	locations, err := pokeapi.ListLocationAreas(&url)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	cfg.nextUrl = locations.Next
	cfg.previousUrl = locations.Previous
	return nil
}

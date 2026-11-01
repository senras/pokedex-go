package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationArea struct {
	Name string `json:"name"`
}

type locationAreasResponse struct {
	Next    string         `json:"next"`
	Previous string        `json:"previous"`
	Results []locationArea `json:"results"`
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"	
	if cfg.nextUrl != "" {
		url = cfg.nextUrl
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data from PokeAPI:", err)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("PokeAPI returned status %s", res.Status)
	}

	var locations locationAreasResponse
	if err := json.NewDecoder(res.Body).Decode(&locations); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	cfg.nextUrl = locations.Next
	cfg.previousUrl = locations.Previous
	return nil
} 

func commandMapb(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.previousUrl != "" {
		url = cfg.previousUrl
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data from PokeAPI:", err)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("PokeAPI returned status %s", res.Status)
	}

	var locations locationAreasResponse
	if err := json.NewDecoder(res.Body).Decode(&locations); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	cfg.nextUrl = locations.Next
	cfg.previousUrl = locations.Previous
	return nil
}
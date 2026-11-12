package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/senras/pokedex-go/internal/pokecache"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func ListPokemonByLocation(location string, cache *pokecache.Cache) ([]string, error) {
	if location == "" {
		return nil, fmt.Errorf("location name cannot be empty")
	}

	if cache == nil {
		return nil, fmt.Errorf("cache cannot be nil")
	}

	locationURL, err := url.JoinPath(LocationAreaURL, location)
	if err != nil {
		return nil, fmt.Errorf("build location URL for %q: %w", location, err)
	}

	if data, ok := cache.Get(locationURL); ok {
		pokemons, err := pokemonNames(data)
		if err != nil {
			return nil, fmt.Errorf("decode cached location %q: %w", location, err)
		}
		return pokemons, nil
	}

	res, err := httpClient.Get(locationURL)
	if err != nil {
		return nil, fmt.Errorf("fetch location %q: %w", location, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("PokeAPI returned status %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read location response: %w", err)
	}

	cache.Add(locationURL, data)

	pokemons, err := pokemonNames(data)
	if err != nil {
		return nil, fmt.Errorf("decode location %q: %w", location, err)
	}
	return pokemons, nil
}

func pokemonNames(data []byte) ([]string, error) {
	var locationArea LocationArea
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return nil, err
	}

	pokemons := make([]string, 0, len(locationArea.PokemonEncounters))
	for _, encounter := range locationArea.PokemonEncounters {
		pokemons = append(pokemons, encounter.Pokemon.Name)
	}
	return pokemons, nil
}

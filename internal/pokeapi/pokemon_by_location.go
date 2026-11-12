package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/senras/pokedex-go/internal/pokecache"
)

func ListPokemonByLocation(location *string, cache *pokecache.Cache) ([]string, error) {

	if location == nil || *location == "" {
		return []string{}, fmt.Errorf("location name cannot be empty")
	}

	if cache == nil {
		return []string{}, fmt.Errorf("cache cannot be nil")
	}

	url := LocationAreaURL + "/" + *location

	if data, ok := cache.Get(url); ok {
		pokemons := []string{}
		if err := json.Unmarshal(data, &pokemons); err != nil {
			return []string{}, err
		}
		return pokemons, nil
	}

	res, err := http.Get(url)

	if err != nil {
		return []string{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []string{}, fmt.Errorf("PokeAPI returned status %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, err
	}
	cache.Add(*location, data)

	pokemons := []string{}
	if err := json.Unmarshal(data, &pokemons); err != nil {
		return []string{}, err
	}

	return pokemons, nil
}

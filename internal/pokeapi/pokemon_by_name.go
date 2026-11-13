package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetPokemonByName(name string) (Pokemon, error) {

	if len(name) == 0 {
		return Pokemon{}, fmt.Errorf("Pokemon name cannot be empty.")

	}

	pokemonURL, err := url.JoinPath(PokemonURL, name)
	if err != nil {
		return Pokemon{}, fmt.Errorf("build Pokemon URL for %q: %w", name, err)
	}

	res, err := httpClient.Get(pokemonURL)

	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("PokeAPI returned status %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil

}

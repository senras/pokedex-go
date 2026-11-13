package pokeapi

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LocationArea struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type LocationAreasResponse struct {
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

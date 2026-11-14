package pokeapi

import (
	"encoding/json"
	"testing"
)

func TestPokemonStatsAndTypesMatchPokeAPIResponse(t *testing.T) {
	const response = `{
		"stats": [
			{"base_stat": 45, "stat": {"name": "hp"}},
			{"base_stat": 49, "stat": {"name": "attack"}}
		],
		"types": [
			{"slot": 1, "type": {"name": "grass"}},
			{"slot": 2, "type": {"name": "poison"}}
		]
	}`

	var pokemon Pokemon
	if err := json.Unmarshal([]byte(response), &pokemon); err != nil {
		t.Fatalf("unmarshal Pokemon response: %v", err)
	}

	if got := pokemon.Stats[0].Stat.Name; got != "hp" {
		t.Errorf("first stat name = %q, want %q", got, "hp")
	}
	if got := pokemon.Stats[0].BaseStat; got != 45 {
		t.Errorf("first stat value = %d, want %d", got, 45)
	}
	if got := pokemon.Types[1].Type.Name; got != "poison" {
		t.Errorf("second type name = %q, want %q", got, "poison")
	}
}

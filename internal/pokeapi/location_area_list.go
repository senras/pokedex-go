package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	if pageURL == nil || *pageURL == "" {
		return LocationAreasResponse{}, fmt.Errorf("location areas URL cannot be empty")
	}

	res, err := http.Get(*pageURL)

	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreasResponse{}, fmt.Errorf("PokeAPI returned status %s", res.Status)
	}

	locations := LocationAreasResponse{}
	if err := json.NewDecoder(res.Body).Decode(&locations); err != nil {
		return LocationAreasResponse{}, err
	}

	return locations, nil
}

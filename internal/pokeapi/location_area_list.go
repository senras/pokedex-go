package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/senras/pokedex-go/internal/pokecache"
)

func ListLocationAreas(pageURL *string, cache *pokecache.Cache) (LocationAreasResponse, error) {
	if pageURL == nil || *pageURL == "" {
		return LocationAreasResponse{}, fmt.Errorf("location areas URL cannot be empty")
	}

	if cache == nil {
		return LocationAreasResponse{}, fmt.Errorf("cache cannot be nil")
	}

	if data, ok := cache.Get(*pageURL); ok {
		locations := LocationAreasResponse{}
		if err := json.Unmarshal(data, &locations); err != nil {
			return LocationAreasResponse{}, err
		}
		return locations, nil
	}

	res, err := http.Get(*pageURL)

	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreasResponse{}, fmt.Errorf("PokeAPI returned status %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	cache.Add(*pageURL, data)

	locations := LocationAreasResponse{}
	if err := json.Unmarshal(data, &locations); err != nil {
		return LocationAreasResponse{}, err
	}
	return locations, nil
}

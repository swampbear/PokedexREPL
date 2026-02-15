package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/swampbear/pokedexcli/internal/models"
)

// used json to Go struct to generate struct

func CommandExplore(conf *Config) error {
	pokeCache, ok := conf.PokeCache.Get(conf.Action)
	if ok {
		locationArea, err := unMarshalLocationAreaFromBytes(pokeCache)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		printEncounters(locationArea)
	} else {
		err := fetchEncounters(conf)
		if err != nil {
			return err
		}
	}

	return nil
}
func unMarshalLocationAreaFromBytes(dat []byte) (models.LocationArea, error) {
	locationArea := models.LocationArea{}
	if err := json.Unmarshal(dat, &locationArea); err != nil {
		return models.LocationArea{}, fmt.Errorf("Failed to unmarshal body: %w", err)
	}
	return locationArea, nil
}

func fetchEncounters(conf *Config) error {

	fullUrl := "https://pokeapi.co/api/v2/location-area/" + conf.Action

	res, err := http.Get(fullUrl)
	if err != nil {
		return fmt.Errorf("Failed to get enconters in location area %w", err)
	}
	defer res.Body.Close()
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Failed to read response body: %w", err)
	}

	locationArea, err := unMarshalLocationAreaFromBytes(dat)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	conf.PokeCache.Add(conf.Action, dat)
	printEncounters(locationArea)
	return nil
}

func printEncounters(locationArea models.LocationArea) {
	fmt.Println("Pokemon available in location-area:")

	for i, encounter := range locationArea.PokemonEncounters {
		fmt.Printf("%v: %s\n", i, encounter.Pokemon.Name)

	}

}

package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// used json to Go struct to generate struct
type LocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

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
func unMarshalLocationAreaFromBytes(dat []byte) (LocationArea, error) {
	locationArea := LocationArea{}
	if err := json.Unmarshal(dat, &locationArea); err != nil {
		return LocationArea{}, fmt.Errorf("Failed to unmarshal body: %w", err)
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

func printEncounters(locationArea LocationArea) {
	fmt.Println("Pokemon available in location-area:")

	for i, encounter := range locationArea.PokemonEncounters {
		fmt.Printf("%v: %s\n", i, encounter.Pokemon.Name)

	}

}

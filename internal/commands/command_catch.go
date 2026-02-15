package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/swampbear/pokedexcli/internal/models"
)

func CommandCatch(conf *Config) error {
	pokemonName := conf.Action
	pokemon, err := fetchPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	caught := tryCatch(pokemon)
	if caught {

		fmt.Printf("%s was caught!\n", pokemon.Name)

		conf.Pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func tryCatch(pokemon models.Pokemon) bool {
	rnd := rand.Intn(500)
	caught := rnd > pokemon.BaseExperience
	return caught
}

func fetchPokemon(pokemonName string) (models.Pokemon, error) {
	fullUrl := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	res, err := http.Get(fullUrl)
	if err != nil {
		return models.Pokemon{}, fmt.Errorf("Failed to get pokemon information: %w", err)
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)

	pmon := models.Pokemon{}

	if err := json.Unmarshal(dat, &pmon); err != nil {
		return models.Pokemon{}, fmt.Errorf("Failed to umarshal pokemon information: %w", err)
	}
	return pmon, nil

}

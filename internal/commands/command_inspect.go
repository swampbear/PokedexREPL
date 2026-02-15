package commands

import (
	"fmt"

	"github.com/swampbear/pokedexcli/internal/models"
)

func CommandInspect(conf *Config) error {
	pokemon, ok := conf.Pokedex[conf.Action]
	if !ok {
		fmt.Println("You have not caught that pokemon")
		return nil
	}
	inspectPokemon(pokemon)
	return nil
}

func inspectPokemon(pokemon models.Pokemon) {

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	stats := pokemon.Stats
	for _, stat := range stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

}

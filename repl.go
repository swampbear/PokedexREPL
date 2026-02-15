package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/swampbear/pokedexcli/internal/commands"
	"github.com/swampbear/pokedexcli/internal/models"
	"github.com/swampbear/pokedexcli/internal/pokecache"
)

var c *commands.Config

func startREPL() {
	// starts scanner and application loop for user interaction

	pokedex := map[string]models.Pokemon{}
	interval := time.Second * 5
	cache := pokecache.NewCache(interval)
	c = &commands.Config{Pokedex: pokedex, PokeCache: cache, Next: "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		words := cleanInput(scanner.Text())
		cmd_txt := words[0]
		// assign second word in cmd_txt to action in the config
		if len(words) > 1 {
			c.Action = words[1]
		}
		cmd, exists := commands.GetCommands()[cmd_txt]
		if exists {
			err := cmd.Callback(c)
			if err != nil {
				fmt.Println(err)
			}

		} else {
			fmt.Println("unknown command")
		}

	}

}

func cleanInput(text string) []string {
	new_text := strings.ToLower(text)
	// .Fields separates strings into a slice of strings separated by spaces
	words := strings.Fields(new_text)
	return words
}

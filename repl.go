package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/swampbear/pokedexcli/internal/commands"
)

var c *commands.Config

func startREPL() {
	// starts scanner and application loop for user interaction

	c = &commands.Config{Next: "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		words := cleanInput(scanner.Text())
		cmd_txt := words[0]
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
	// Fields separates strings into a slice of strings separated by spaces
	words := strings.Fields(new_text)
	return words
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	//bufio scanner for reading input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		words := cleanInput(scanner.Text())
		cmd_txt := words[0]
		cmd, exists := getCommands()[cmd_txt]
		if exists {
			cmd.callback()
		} else {
			fmt.Println("unknown command")
		}

	}

}

func cleanInput(text string) []string {
	new_text := strings.ToLower(text)
	words := strings.Fields(new_text)
	return words
}

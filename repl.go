package main

import (
	"strings"
)

func cleanInput(text string) []string {
	new_text := strings.ToLower(text)
	words := strings.Fields(new_text)
	return words
}

package main

import (
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	for i, word := range words {
		words[i] = strings.Trim(word, ".,!?;:")
	}
	return words
}
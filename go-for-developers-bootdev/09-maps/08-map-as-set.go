package main

import (
	"fmt"
	"strings"
)

func countDistinctWords(messages []string) int {
	words := make(map[string]struct{})
	for _, message := range messages {
		for _, word := range strings.Fields(message) {
			words[strings.ToLower(word)] = struct{}{}
		}
	}

	return len(words)
}

func main() {
	messages := []string{"hello", "Hello", "HELLO", "world", "World", "WORLD"}
	distinctWords := countDistinctWords(messages)
	fmt.Println(distinctWords)
}

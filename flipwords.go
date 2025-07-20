package main

import (
	"fmt"
	"strings"
)

func moveLastWordToFront(s string) string {
	words := strings.Fields(s)
	n := len(words)
	if n == 0 {
		return s
	}
	last := words[n-1]
	rest := words[:n-1]
	newWords := append([]string{last}, rest...)
	return strings.Join(newWords, " ")
}

func main() {
	result := "Hello to the world"
	fmt.Println(moveLastWordToFront(result)) // Output: "world Hello to the"
}

package main

import "fmt"

// Convert a lowercase letter to uppercase
func toupperByte(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}

// Convert an uppercase letter to lowercase
func tolowerByte(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

// Check if character is a separator
func isoperator(c byte) bool {
	return c == ' ' || c == '-' || c == '_'
}

// Convert string to PascalCase
func topascal(s string) string {
	result := ""
	captalizeNext := true

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if isoperator(ch) {
			captalizeNext = true
			continue
		}

		if captalizeNext {
			result += string(toupperByte(ch))
			captalizeNext = false
		} else {
			result += string(tolowerByte(ch))
		}
	}
	return result
}

func main() {
	fmt.Println(topascal("hello_world"))        // HelloWorld
	fmt.Println(topascal("make-it_pascal case")) // MakeItPascalCase
	fmt.Println(topascal("GoLang"))              // Golang
}

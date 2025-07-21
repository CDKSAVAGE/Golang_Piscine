package main

import "fmt"

func Atoi(s string) int {
	result := 0
	sign := 1
	start := 0

	// Handle optional sign
	if len(s) > 0 && s[0] == '-' {
		sign = -1
		start = 1
	} else if len(s) > 0 && s[0] == '+' {
		start = 1
	}

	// Loop through each character
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break // stop if not a digit
		}
		result = result*10 + int(s[i]-'0')
	}

	return result * sign
}

func main() {
	fmt.Println(Atoi("123"))     // 123
	fmt.Println(Atoi("-456"))    // -456
	fmt.Println(Atoi("+789"))    // 789
	fmt.Println(Atoi("42abc"))   // 42 (stops at non-digit)
}

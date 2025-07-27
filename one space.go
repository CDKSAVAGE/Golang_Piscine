package main

import "fmt"

func trimFields(s string) string {
	start := 0
	end := len(s) - 1
	for start <= end && s[start] == ' ' {
		start++
	}
	for end >= start && s[end] == ' ' {
		end--
	}
	return s[start : end+1]
}

func addOneSpace(s string) string {
	word := trimFields(s)
	result := ""
	space := false
	for i := 0; i < len(word); i++ {
		if word[i] != ' ' {
			result += string(word[i])
			space = false
		} else if !space {
			// only add one space
			result += " "
			space = true
		}
	}
	return result
}

func main() {
	fmt.Println(addOneSpace("   Hello   World     I   ")) // "Hello World I"
}

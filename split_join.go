package main

import "fmt"

func splitword(s string) []string {
    var words []string
    word := ""
    for i := 0; i < len(s); i++ {
        if s[i] != ' ' {
            word += string(s[i])
        } else if word != "" {
            words = append(words, word)
            word = ""
        }
    }
    if word != "" {
        words = append(words, word)
    }
    return words
}

func joinword(words []string) string {
    result := ""
    for i := 0; i < len(words); i++ {
        result += words[i]
        if i != len(words)-1 {
            result += " " // Add space between words
        }
    }
    return result
}

func main() {
    input := "Hello World hb"

    // Split into words
    splitResult := splitword(input)
    fmt.Println("Split:", splitResult) // Output: [Hello World hb]

    // Join words back into a sentence
    joinedResult := joinword(splitResult)
    fmt.Println("Joined:", joinedResult) // Output: Hello World hb
} 

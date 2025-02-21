package main

import "fmt"

func main() {
	for i := 'a'; i <= 'z'; i++ {
		if (i-'a')%2 == 0 { // Odd index (starting from 0)
			fmt.Printf("%c", i) // Lowercase for odd index
		} else { // Even index
			fmt.Printf("%c", i-32) // Convert to uppercase by subtracting 32 (ASCII difference)
		}
	}
	fmt.Println() // Print a newline at the end
}

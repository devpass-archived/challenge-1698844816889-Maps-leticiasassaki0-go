package main

import (
	"fmt"
)

func countFruits(fruits []string) map[string]int {
	// Add your solution here!
	return nil
}

func main() {
	fruits := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	fruitCount := countFruits(fruits)

	for fruit, count := range fruitCount {
		fmt.Printf("%s: %d\n", fruit, count)
	}
}
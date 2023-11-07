package main

import (
	"fmt"
)

func countFruits(fruits []string) map[string]int {
	return 0
}

func main() {
	fruits := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	fruitCount := countFruits(fruits)

	for fruit, count := range fruitCount {
		fmt.Printf("%s: %d\n", fruit, count)
	}
}

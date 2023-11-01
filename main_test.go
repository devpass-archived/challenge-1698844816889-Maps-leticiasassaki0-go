package main

import (
	"reflect"
	"testing"
)

func TestCountFruits(t *testing.T) {
	fruits := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	fruitCount := countFruits(fruits)

	expected := map[string]int{"apple": 3, "banana": 2, "orange": 1}

	if !reflect.DeepEqual(fruitCount, expected) {
		t.Errorf("Expected %v, but got %v", expected, fruitCount)
	}
}
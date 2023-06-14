package main

import (
	"fmt"
	"sort"
)

func ordenarStrings(slice []string) []string {
	if len(slice) == 0 {
		panic("O slice não pode estar vazio.")
	}
	sortedSlice := make([]string, len(slice))
	copy(sortedSlice, slice)
	sort.Strings(sortedSlice)
	return sortedSlice
}

func main() {
	strings := []string{"banana", "abacaxi", "laranja"}
	fmt.Println(ordenarStrings(strings)) // Output: [abacaxi banana laranja]
}

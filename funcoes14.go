package main

import (
	"fmt"
)

func numerosComuns(slice1, slice2 []int) []int {
	if len(slice1) == 0 || len(slice2) == 0 {
		panic("Ambos os slices devem conter elementos.")
	}
	numerosComuns := []int{}
	numerosMap := make(map[int]bool)
	for _, num := range slice1 {
		numerosMap[num] = true
	}
	for _, num := range slice2 {
		if numerosMap[num] {
			numerosComuns = append(numerosComuns, num)
		}
	}
	return numerosComuns
}

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4, 5, 6}
	fmt.Println(numerosComuns(slice1, slice2)) // Output: [3 4]
}

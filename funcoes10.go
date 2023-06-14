package main

import (
	"errors"
	"fmt"
)

func sortIntSlice(numbers []int) ([]int, error) {
	if len(numbers) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	sorted := make([]int, len(numbers))
	copy(sorted, numbers)

	for i := 0; i < len(sorted)-1; i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j] < sorted[i] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	return sorted, nil
}

func main() {
	intSlice := []int{9, 5, 7, 1, 3}
	result, err := sortIntSlice(intSlice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println(result)
	}
}

package main

import (
	"errors"
	"fmt"
)

func filterEvenNumbers(numbers []int) ([]int, error) {
	if len(numbers) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	result := make([]int, 0)
	for _, num := range numbers {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result, nil
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result, err := filterEvenNumbers(intSlice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println(result)
	}
}

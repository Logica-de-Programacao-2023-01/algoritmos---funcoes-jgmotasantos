package main

import (
	"errors"
	"fmt"
)

func applyFunctionToInts(numbers []int, fn func(int) int) ([]int, error) {
	if len(numbers) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = fn(num)
	}
	return result, nil
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	doubleFn := func(num int) int {
		return num * 2
	}
	result, err := applyFunctionToInts(intSlice, doubleFn)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println(result)
	}
}

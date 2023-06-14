package main

import (
	"fmt"
)

func numeroPresente(numero int, sliceInteiros []int) bool {
	if len(sliceInteiros) == 0 {
		panic("O slice não pode estar vazio.")
	}
	for _, num := range sliceInteiros {
		if num == numero {
			return true
		}
	}
	return false
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	numero := 3
	fmt.Println(numeroPresente(numero, slice)) // Output: true
}

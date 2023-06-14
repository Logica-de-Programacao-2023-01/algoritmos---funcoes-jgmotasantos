package main

import (
	"fmt"
)

func encontrarSegundoMaiorValor(slice []int) int {
	max := slice[0]
	segundoMax := slice[0]
	for _, valor := range slice {
		if valor > max {
			segundoMax = max
			max = valor
		} else if valor > segundoMax && valor < max {
			segundoMax = valor
		}
	}
	return segundoMax
}

func main() {
	sliceInteiros := []int{1, 3, 2, 5, 4}
	segundoMaior := encontrarSegundoMaiorValor(sliceInteiros)
	fmt.Println("Segundo maior valor:", segundoMaior)
}

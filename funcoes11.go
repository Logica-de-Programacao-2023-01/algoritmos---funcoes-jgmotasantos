package main

import (
	"fmt"
	"math"
)

func ehPrimo(numero int) bool {
	if numero < 0 {
		panic("O número deve ser positivo.")
	}
	if numero < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(numero))); i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	numero := 17
	fmt.Println(ehPrimo(numero)) // Output: true
}

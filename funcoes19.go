package main

import (
	"fmt"
)

func obterPrimos(numero int) []int {
	if numero < 0 {
		panic("O número deve ser positivo.")
	}
	primos := []int{}
	for i := 2; i <= numero; i++ {
		if isPrimo(i) {
			primos = append(primos, i)
		}
	}
	return primos
}

func isPrimo(numero int) bool {
	if numero < 2 {
		return false
	}
	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	numero := 20
	fmt.Println(obterPrimos(numero)) // Output: [2 3 5 7 11 13 17 19]
}

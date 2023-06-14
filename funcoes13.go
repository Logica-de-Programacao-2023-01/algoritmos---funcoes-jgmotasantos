package main

import (
	"fmt"
	"strconv"
)

func somaDigitos(numero int) int {
	if numero < 0 {
		panic("O número deve ser positivo.")
	}
	digitos := strconv.Itoa(numero)
	soma := 0
	for _, digito := range digitos {
		digitoInt, _ := strconv.Atoi(string(digito))
		soma += digitoInt
	}
	return soma
}

func main() {
	numero := 12345
	fmt.Println(somaDigitos(numero)) // Output: 15
}

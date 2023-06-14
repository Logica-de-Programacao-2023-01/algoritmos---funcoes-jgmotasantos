package main

import (
	"fmt"
)

func aplicarFuncao(slice []int, funcao func(int) int) int {
	if len(slice) == 0 {
		panic("O slice não pode estar vazio.")
	}
	soma := 0
	for _, num := range slice {
		resultado := funcao(num)
		soma += resultado
	}
	return soma
}

func dobrarNumero(numero int) int {
	return numero * 2
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(aplicarFuncao(slice, dobrarNumero)) // Output: 30 (1*2 + 2*2 + 3*2 + 4*2 + 5*2)
}

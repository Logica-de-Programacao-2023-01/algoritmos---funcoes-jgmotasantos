package main

import (
	"fmt"
)

func encontrarPosicao(slice []int, valor int) int {
	for i, v := range slice {
		if v == valor {
			return i
		}
	}
	return -1
}

func main() {
	sliceInteiros := []int{1, 3, 2, 5, 4}
	valor := 2
	posicao := encontrarPosicao(sliceInteiros, valor)
	fmt.Println("Posição do valor", valor, ":", posicao)
}

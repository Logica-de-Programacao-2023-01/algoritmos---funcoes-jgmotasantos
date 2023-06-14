package main

import (
	"fmt"
	"strings"
)

func contarVogais(str string) int {
	vogais := []string{"a", "e", "i", "o", "u"}
	count := 0
	for _, char := range strings.ToLower(str) {
		if contains(vogais, string(char)) {
			count++
		}
	}
	return count
}

func contains(slice []string, valor string) bool {
	for _, item := range slice {
		if item == valor {
			return true
		}
	}
	return false
}

func main() {
	str := "Hello, World!"
	qtdVogais := contarVogais(str)
	fmt.Println("Quantidade de vogais:", qtdVogais)
}

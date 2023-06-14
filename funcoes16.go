package main

import (
	"fmt"
	"strings"
)

func substituirVogais(str string) string {
	if len(str) == 0 {
		panic("A string não pode estar vazia.")
	}
	vogais := "aeiouAEIOU"
	resultado := strings.Builder{}
	for _, char := range str {
		if strings.ContainsRune(vogais, char) {
			resultado.WriteRune('*')
		} else {
			resultado.WriteRune(char)
		}
	}
	return resultado.String()
}

func main() {
	str := "Hello, World!"
	fmt.Println(substituirVogais(str)) // Output: H*ll*, W*rld!
}

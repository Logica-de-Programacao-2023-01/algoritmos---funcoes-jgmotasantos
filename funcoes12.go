package main

import (
	"fmt"
	"unicode"
)

func filtrarMaiusculas(strings []string) []string {
	if len(strings) == 0 {
		panic("O slice não pode estar vazio.")
	}
	resultado := []string{}
	for _, str := range strings {
		if len(str) > 0 && unicode.IsUpper([]rune(str)[0]) {
			resultado = append(resultado, str)
		}
	}
	return resultado
}

func main() {
	strings := []string{"Abacaxi", "banana", "Cenoura"}
	fmt.Println(filtrarMaiusculas(strings)) // Output: [Abacaxi Cenoura]
}

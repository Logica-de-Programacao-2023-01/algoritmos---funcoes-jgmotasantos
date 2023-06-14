package main

import (
	"fmt"
)

func filtrarStrings(slice []string) []string {
	if len(slice) == 0 {
		panic("O slice não pode estar vazio.")
	}
	resultado := []string{}
	for _, str := range slice {
		if len(str) > 5 {
			resultado = append(resultado, str)
		}
	}
	return resultado
}

func main() {
	slice := []string{"banana", "abacaxi", "laranja", "maçã", "uva"}
	fmt.Println(filtrarStrings(slice))
}

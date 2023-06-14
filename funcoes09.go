package main

import (
	"errors"
	"fmt"
	"strings"
)

func extractWordsFromString(str string) ([]string, error) {
	if str == "" {
		return nil, errors.New("A string está vazia")
	}

	words := strings.Fields(str)
	return words, nil
}

func main() {
	text := "Olá, mundo! Bem-vindo ao Go."
	result, err := extractWordsFromString(text)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println(result)
	}
}

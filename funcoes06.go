package main

import (
	"errors"
	"fmt"
)

func concatenateStrings(strings []string) (string, error) {
	if len(strings) == 0 {
		return "", errors.New("O slice está vazio")
	}

	result := ""
	for i, str := range strings {
		if i > 0 {
			result += ", "
		}
		result += str
	}

	return result, nil
}

func main() {
	strSlice := []string{"Hello", "world", "Gopher"}
	concatenated, err := concatenateStrings(strSlice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println(concatenated)
	}
}

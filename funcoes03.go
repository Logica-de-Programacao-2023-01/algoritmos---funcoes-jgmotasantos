package main

import (
	"fmt"
	"strings"
)

func concatenarStrings(slice []string) string {
	return strings.Join(slice, "")
}

func main() {
	sliceStrings := []string{"Hello", " ", "World!"}
	concatenacao := concatenarStrings(sliceStrings)
	fmt.Println("Concatenação:", concatenacao)
}

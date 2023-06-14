package main

import "fmt"

func calcularMedia(slice []int) float64 {
	total := 0
	for _, valor := range slice {
		total += valor
	}
	return float64(total) / float64(len(slice))
}

func main() {
	sliceInteiros := []int{1, 2, 3, 4, 5}
	media := calcularMedia(sliceInteiros)
	fmt.Println("Média:", media)
}

package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println(numeros)

	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func escrever(texto string, numeros ...int) {
	fmt.Println(texto, numeros)
	for _, numero := range numeros {
		fmt.Println(numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5)
	fmt.Println(totalDaSoma)

	escrever("Olá", 1, 2, 3, 4, 5)
}

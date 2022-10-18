package main

import "fmt"

func main() {
	soma := somar(2, 6)
	fmt.Println(soma)

	var f = func(msg string) string {
		fmt.Println(msg)
		return msg
	}

	resultado := f("Texto da func anônima")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubstracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubstracao)
}

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

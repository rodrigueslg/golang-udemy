package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Executou o recover")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	panic("A MÉDIA FOI CALCULADA")
	return media >= 6
}

func main() {
	fmt.Println(alunoEstaAprovado(8, 6))
	fmt.Println("Pós execução da função")
}

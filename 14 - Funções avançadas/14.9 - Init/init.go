package main

import "fmt"

var n int

func init() {
	n = 10
	fmt.Println("Executando a função init")
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}

package main

import "fmt"

func main() {

	//VALUE
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//REFERENCE
	var variavel3 int = 100
	var ponteiro *int = &variavel3
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
}

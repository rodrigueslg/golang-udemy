package main

import "fmt"

func main() {
	// ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String"
	fmt.Println(variavel1, variavel2)

	//RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//UNÁRIOS
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 20
	numero *= 3
	fmt.Println(numero)
}

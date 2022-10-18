package main

import (
	"errors"
	"fmt"
)

func main() {

	// INTEIROS
	var numero int64 = -1000000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// REAIS (FLOAT)
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// STRING
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// BOOLEAN
	var bool1 bool = true
	fmt.Println(bool1)

	// ERROR
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}

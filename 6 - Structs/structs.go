package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Structs")

	var u usuario
	u.nome = "João"
	u.idade = 25
	u.endereco.logradouro = "Rua dos Bobos"
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}
	u2 := usuario{"Maria", 30, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{idade: 35}
	fmt.Println(u3)
}

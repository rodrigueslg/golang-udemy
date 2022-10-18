package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second / 4)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second / 4)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, letra := range "Palavra" {
		fmt.Println(string(letra))
	}

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}

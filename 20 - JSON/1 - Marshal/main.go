package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade int    `json:"idade"`
}

func main() {
	c := Cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJson)
	fmt.Println(bytes.NewBuffer(cachorroJson))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorroJson2, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJson2)
	fmt.Println(bytes.NewBuffer(cachorroJson2))
}

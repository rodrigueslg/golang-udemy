package main

import (
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
	cachorroJson := `{"nome":"Rex","raca":"Dalmata","idade":3}`
	var c Cachorro
	if erro := json.Unmarshal([]byte(cachorroJson), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	cachorroJson2 := `{"nome":"Toby","raca":"Poodle"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorroJson2), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)
}

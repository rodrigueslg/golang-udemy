package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Silva",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Londrina",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"])

	delete(usuario2, "curso")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)
}

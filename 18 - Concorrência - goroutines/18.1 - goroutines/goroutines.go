package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	// Concorrência é quando duas ou mais tarefas podem começar, executar e terminar em tempos diferentes.
	// Paralelismo é quando duas ou mais tarefas são executadas simultaneamente.

	go escrever("Olá mundo!")
	go escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}

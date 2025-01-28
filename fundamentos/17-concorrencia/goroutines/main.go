package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrência != Paralelismo
	go escrever("Olá mundo") // go routine
	escrever("Programando em go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

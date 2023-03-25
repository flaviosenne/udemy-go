package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		i++
		fmt.Println("incrementando i: ", i)
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("incrementando j: ", j)
		time.Sleep(time.Second)
	}

	nomes := []string{"João", "José", "Ana"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "joao",
		"sobrenome": "senne",
	}

	for indice, valor := range usuario {
		fmt.Println(indice, valor)
	}

	// lopp infinito
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}

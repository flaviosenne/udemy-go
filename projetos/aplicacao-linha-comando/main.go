package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	apicacao := app.Gerar()
	if erro := apicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}

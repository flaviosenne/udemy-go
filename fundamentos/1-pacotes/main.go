package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("escrevendo do main")
	auxiliar.Escrever()

	err := checkmail.ValidateFormat("joao@gmail.com")
	fmt.Println(err)
}

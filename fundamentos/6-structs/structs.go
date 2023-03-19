package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int8
}

func main() {
	var u usuario
	fmt.Println(u)
	u.nome = "joao"
	u.idade = 23
	fmt.Println(u)

	enderecoExemplo := endereco{"rua 1", 0}

	u2 := usuario{"Jose", 23, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{nome: "Ana"}
	fmt.Println(u3)

}

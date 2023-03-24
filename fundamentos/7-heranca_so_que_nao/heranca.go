package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    // herança em go
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	p1 := pessoa{
		nome:      "nome",
		sobrenome: "sobrenome",
		idade:     12,
		altura:    180,
	}

	fmt.Println(p1)

	e1 := estudante{p1, "aaa", "11111"}
	fmt.Println(e1.nome)
}

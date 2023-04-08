package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
}

func (u usuario) salvar() {
	fmt.Printf("salvar usuario %s no banco\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"joao", 10}
	usuario2 := usuario{"davi", 18}
	usuario1.salvar()
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println("maior de idade ", maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println("idade do usuario 2: ", usuario2.idade)
}

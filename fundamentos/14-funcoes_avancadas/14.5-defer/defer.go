package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoEstaAprovado(n1, n2 int) bool {
	// ao invés de colocar esse log no if e fora do if usando o defer, ele sempre será usado no final da função
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Enrando na função para verificar se o aluno foi aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		// fmt.Println("Média calculada. Resultado será retornado")
		return true
	}

	// fmt.Println("Média calculada. Resultado será retornado")
	return false
}

func main() {
	defer funcao1()
	// defer = adiar
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}

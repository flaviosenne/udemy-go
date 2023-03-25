package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("maior q 15")
	} else {
		fmt.Println("menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("numero maior q 0")
		fmt.Println(outroNumero, "aqui é acessível")
	} else if numero < -10 {
		fmt.Println(outroNumero, "aqui é acessível")
	} else {
		fmt.Println(outroNumero, "entre 0 e -10")

	}
	fmt.Println("outroNumero aqui não é acessível")

}

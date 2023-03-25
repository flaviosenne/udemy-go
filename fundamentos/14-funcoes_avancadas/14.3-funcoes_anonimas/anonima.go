package main

import "fmt"

func main() {
	func() {
		fmt.Println("olá mundo")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("passando parametro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("passando parametro")

	fmt.Println(retorno)
}

package main

import "fmt"

func main() {
	// causa deadlock
	// canal := make(chan string)
	// canal <- "olá mundoo"

	// mensagem := <-canal
	// fmt.Println(mensagem)

	canalComBuffer := make(chan string, 2)
	canalComBuffer <- "olá mundoo 2"
	canalComBuffer <- "olá mundoo 3"

	mensagem2 := <-canalComBuffer
	mensagem3 := <-canalComBuffer
	fmt.Println(mensagem2, mensagem3)

}

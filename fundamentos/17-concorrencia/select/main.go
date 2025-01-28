package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		// aqui existe uma perd de tempo por parte do canal 1, pois ele tem q esperar
		// 2 segundos do canal 2 para poder pegar a mensagem do seu próprio canal
		// mensagemCanal1 := <-canal1
		// fmt.Println(mensagemCanal1)

		// mensagemCanal2 := <-canal2
		// fmt.Println(mensagemCanal2)

		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}

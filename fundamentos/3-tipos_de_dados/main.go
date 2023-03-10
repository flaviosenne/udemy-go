package main

import (
	"errors"
	"fmt"
)

func main() {
	// numeros inteiros
	// int8, int16, int32, int64
	// o tipo int usa como referencia a arquitetura do
	// computador ex: 32bits / 64bits
	numero := 1000
	fmt.Println(numero)

	// uint suporta apenas valores positivos
	var numero2 uint = 0
	fmt.Println(numero2)

	// alias para um tipo do int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	// alias para um tipo do int8 = byte
	var numero4 byte = 127
	fmt.Println(numero4)

	// numeros reais float32, float64

	var numeroReal1 float32 = 123.21
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12331409890.9021
	fmt.Println(numeroReal2)

	numeroReal3 := 1235.354
	fmt.Println(numeroReal3)

	// tipo string
	var str string = "text"
	fmt.Println(str)

	str2 := "text2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	var text string
	var num int
	fmt.Println(text)
	fmt.Println(num)

	var booleano1 bool = true
	var booleano2 bool = false
	var booleano3 bool
	fmt.Println(booleano1, booleano2, booleano3)

	var erro1 error
	var erro2 error = errors.New("erro interno")
	fmt.Println(erro1, erro2)

}

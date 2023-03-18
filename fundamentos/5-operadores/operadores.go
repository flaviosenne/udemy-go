package main

import "fmt"

func main() {
	// aritiméticos
	fmt.Println("-----------------")
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 2
	restoDivisao := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25

	var soma2 = numero1 + int16(numero2)
	fmt.Println(soma2)

	// atribuição
	fmt.Println("-----------------")
	var variavel1 string = "String"
	variavel2 := "String por inferencia"
	fmt.Println(variavel1, variavel2)

	// operadores relacionais
	fmt.Println("-----------------")
	fmt.Println(1 > 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// operadores lógicos
	fmt.Println("-----------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// operadores unários
	fmt.Println("-----------------")
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 10
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 10
	fmt.Println(numero)
	numero *= 3
	fmt.Println(numero)
	numero /= 3
	fmt.Println(numero)
	numero %= 3
	fmt.Println(numero)

	// operador ternario não existe
	// texto := numero > 5 ? "maior que 5" < "menor que 5"
}

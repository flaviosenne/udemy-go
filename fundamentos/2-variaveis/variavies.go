package main

import "fmt"

func main() {
	var variavel1 string = "Varievel 1"
	variavel2 := "Varievel 2" // inferência de tipo
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "bababa"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}

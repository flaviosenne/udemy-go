package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("Uma string")
	generica(1)
	generica(true)

	fmt.Println(1, true, "String", float64(12))

	mapa := map[interface{}]interface{}{
		1:        "123",
		"String": 1243,
		true:     float64(10),
	}

	fmt.Println(mapa)
}

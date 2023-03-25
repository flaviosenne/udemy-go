package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Joao",
		"sobrenome": "Senne",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "Senne",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "curso")
	fmt.Println(usuario2)
	usuario2["signo"] = map[string]string{
		"nome": "Ares",
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["primeiro"])
}

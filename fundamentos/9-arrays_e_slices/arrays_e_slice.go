package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slice")

	var array1 [5]int
	fmt.Println(array1)
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array3), reflect.TypeOf(slice))

	slice = append(slice, 14, 15)
	fmt.Println(slice)

	slice2 := array2[1:3] // fatia o array apontando para o endereço,
	//caso haja alteração no array, vai refletir na variavel
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

}

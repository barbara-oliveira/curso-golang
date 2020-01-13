package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array --  Tamanho fixo
	s1 := []int{1, 2, 3}  //slice  -- Tamanho variado

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//Slice não é um array! Slice define o pedaço de um array somente aponta os valores do array
	s2 := a2[1:3] // Vai do índice 1 ao 3 (não incluindo o índice 3) do array
	fmt.Println(a2, s2)

	s3 := a2[:2] //nova slice que aponta para o mesmo array a2
	fmt.Println(s3)

	//vc pode imaginar o slice como: tamanho e um ponteiro para um elemento de um array

	s4 := s2[:1] //slice de um slice
	fmt.Println(s2, s4)

}

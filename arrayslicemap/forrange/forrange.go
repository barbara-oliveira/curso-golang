package main

import "fmt"

func main() {
	numeros := [...]int{100, 21, 0, 5, 9} //compilador conta o tamanho do array

	for i, numero := range numeros { //é como um foreach
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { //Sem usar o índice
		fmt.Println(num)
	}

}

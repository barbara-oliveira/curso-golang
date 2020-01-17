package main

import "fmt"

var soma = func(a, b int) int { //Função anônima
	return a + b
}

func main() {
	fmt.Println("Soma: ", soma(1, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Printf("Sub: %d", sub(89, 9))
}

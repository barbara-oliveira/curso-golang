package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{} // variável do tipo genérico
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{} //declaração de um tipo genérico
	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang"}
	fmt.Println(coisa2)
}

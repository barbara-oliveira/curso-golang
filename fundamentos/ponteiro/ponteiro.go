package main

import "fmt"

func main() {
	i := 1

	//Go não tem aritmética de ponteiro
	var p *int = nil

	p = &i //pegando o endereço da variável i e atribuindo ao ponteiro
	*p++   //Pegou o valor associado ao ponteiro e incrementou o valor
	i++

	fmt.Println(p, *p, i) // p é o endereço de memória
}

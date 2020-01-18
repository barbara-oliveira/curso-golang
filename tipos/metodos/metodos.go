package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) { //ponteiro são necessários
	partes := strings.Split(nomeCompleto, " ") //para alterar os valoes internos dos tipos
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Julio", "Cesar"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Cecília Marcon")
	fmt.Println(p1.getNomeCompleto())
}

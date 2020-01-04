package main

import (
	"fmt"
	m "math" //m está referenciando o pacote math
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo float64 inferido pelo compilador

	//forma reduzida de criar uma variavel
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é ", area)

	//declarando uma lista de contantes
	const (
		a = 1
		b = 2
	)

	//declarando uma lista de variáveis
	var (
		c = 2
		d = 1
	)

	fmt.Println(a, b, c, d)

	var e, f, g, h bool = true, false, true, true
	fmt.Println(e, f, g, h)

	i, j, k := 2, false, "Oieeee" //Os tipos da variável não mudam ao longo do programa
	fmt.Println(i, j, k)
}

/*
Obs: Na linguagem Go, se vc nãoo utilizar a variável que foi declarada,
vai ocorrer um erro de compilação e o programa não irá rodar
*/

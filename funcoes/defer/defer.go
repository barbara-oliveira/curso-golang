package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim!") // defer = executa no último momento possível da sentença do método
	if numero >= 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(5000))
}

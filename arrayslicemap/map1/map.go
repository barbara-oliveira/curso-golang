package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// Maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[23476513] = "Antônio"
	aprovados[45456575] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[23476513])
	delete(aprovados, 45456575)
	fmt.Println(aprovados)
}

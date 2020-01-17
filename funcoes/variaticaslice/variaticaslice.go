package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	//somente variáveis do tipo slice são permitidas para serem inseridas
	//como parâmetros de funções variáticas. (Não funciona com arrays)

	aprovados := []string{"Julie", "Luke", "Miucha", "Bob"}
	imprimirAprovados(aprovados...)
}

package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José da Silva": 1234.23,
		"Maria Santos":  3000,
		"Pedro Paulo":   5250.60,
	}

	funcsESalarios["Luana Lua"] = 12000.50
	delete(funcsESalarios, "Inexistente")
	fmt.Println(funcsESalarios)

	for nome, salario := range funcsESalarios { // o primeiro valor do for sempre será achave do map
		fmt.Println(nome, salario)
	}

}

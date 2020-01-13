package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 1342.56,
			"Guga Braga":     4563.67,
		},
		"J": {
			"José João": 1200.70,
		},
		"P": {
			"Paulo": 3456.23,
		},
	}

	fmt.Println(funcsPorLetra)
	delete(funcsPorLetra, "P")
	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}

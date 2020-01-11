package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch {
	case nota == 10:
		fallthrough
	case nota == 9:
		return "A"
	case nota == 8 || nota == 7:
		return "B"
	case nota == 5 || nota == 6:
		return "C"
	case nota == 4 || nota == 3:
		return "D"
	case nota == 2 || nota == 1 || nota == 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(100))
	fmt.Println(notaParaConceito(9.1))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(5.7))
}

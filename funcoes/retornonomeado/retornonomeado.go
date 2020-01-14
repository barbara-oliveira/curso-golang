package main

import "fmt"

func troca(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo
}

func main() {
	fmt.Println(troca(1, 2))

	segundo, primeiro := troca(3, 4)
	fmt.Println(segundo, primeiro)
}

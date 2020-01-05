package main

import "fmt"

func main() {
	fmt.Print("Mesma ")
	fmt.Print("Linha")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	//fmt.Println("O valor de x é " + x) não funciona

	xs := fmt.Sprint(x) //converte para string
	fmt.Println("O valor de x é " + xs)

	fmt.Println("O valor de x é", x, "oieee") // sem concatenação

	fmt.Printf("O valor de x é %.2f", x) //print formatado

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d)

}

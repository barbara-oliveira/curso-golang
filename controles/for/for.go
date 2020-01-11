package main

import "fmt"

func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	//for { } laço infinito
}

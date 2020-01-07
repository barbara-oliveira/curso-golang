package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings: ", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 4)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 2 <= 3)
	fmt.Println(">=", 3 >= 3)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	//Somente os valores estão sendo levados em consideração
	p1 := Pessoa{"Bilbo"}
	p2 := Pessoa{"Bilbo"}

	fmt.Println("Structs:", p1 == p2)
}

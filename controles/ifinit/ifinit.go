package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { //Tbm suportado no switch
		fmt.Println("Ganhou!!")
		fmt.Println(i) // A variável i só fica definida dentro dos blocos if/else
	} else {
		fmt.Println("Perdeu!!!")
	}
}

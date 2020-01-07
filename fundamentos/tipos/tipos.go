package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//números inteiros
	fmt.Println(1, 2, 10000)
	fmt.Println("Literal inteiro é ", reflect.TypeOf(32000))

	//sem sinal só positivos... uint8 uint16 uint16 uint64
	var b byte = 255 //byte é um alias para uint8
	fmt.Println("Byte é ", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64 //maior valor do tipo int64
	fmt.Println("O valor máximo do int é ", i1)
	fmt.Println("O tipo de i1 é ", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa ummapeamento da tabela unicode
	fmt.Println("O rune é ", reflect.TypeOf(i2))
	fmt.Println(i2)

	//números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é ", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é ", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("A variável é do tipo", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//String
	s1 := "Olá meu nome é Bárbara"
	fmt.Println(s1, "!")
	fmt.Println(" O tamanho de s1 é", len(s1))

	//String com múltiplas linhas
	s2 := `Oieeee
	ôôô vida de gado
	Pôvo Largado 
	Pôvo feliz`

	fmt.Println(s2)
	fmt.Println("O tamanho de s2 é", len(s2))

	//Não existe o tipo char no Go!
	char := 'a'
	fmt.Println("o tipo de char é", reflect.TypeOf(char))

}

package main

import "fmt"

type item struct {
	produtoId int
	qtde      int
	preco     float64
}

type pedido struct {
	userId int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido1 := pedido{
		userId: 1,
		itens: []item{
			item{1, 2, 12.7},
			item{2, 1, 23.49},
			item{11, 100, 3.13},
		},
	}

	fmt.Printf("O valor total do pedido é %.2f", pedido1.valorTotal())
}

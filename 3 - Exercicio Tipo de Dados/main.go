package main

import "fmt"

type Produto struct {
	nome  string
	preco float64
}

func main() {
	var p Produto
	p.nome = "Relógio"
	p.preco = 259.50

	fmt.Printf("O %s custa R$ %f", p.nome, p.preco)

}

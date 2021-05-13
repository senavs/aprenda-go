package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type carros []carro

func (c carros) Len() int {
	return len(c)
}

func (c carros) Less(i, j int) bool {
	return c[i].potencia < c[j].potencia
}

func (c carros) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {
	c := []carro{
		{"Porsche", 300, 5},
		{"Chevrole", 50, 5},
		{"Fusca", 500, 30},
	}

	cs := carros(c)

	fmt.Println(cs)
	sort.Sort(cs)
	fmt.Println(cs)
}

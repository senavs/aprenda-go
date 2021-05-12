package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) saudacao() {
	fmt.Println(p.nome, "diz bom dia!")
}

func main() {
	p1 := pessoa{"Mauricio", 30}
	p1.saudacao()

	// saudacao()  // error
	// saudacao(p1)()  // error
}

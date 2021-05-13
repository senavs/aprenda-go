package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
	Nome           string
	Sobrenome      string
	Idade          int
	Profissao      string
	ContanBancaria float64
}

func main() {
	p1 := pessoa{
		Nome:           "James",
		Sobrenome:      "Bond",
		Idade:          40,
		Profissao:      "Agente Secreto",
		ContanBancaria: 1000000.50,
	}

	p2 := pessoa{
		Nome:           "Anakin",
		Sobrenome:      "Skywalker",
		Idade:          60,
		Profissao:      "Vilão",
		ContanBancaria: 50000000000.88,
	}

	j1, err1 := json.Marshal(p1)
	if err1 != nil {
		fmt.Println("error:", err1)
	}

	j2, err2 := json.Marshal(p2)
	if err2 != nil {
		fmt.Println("error:", err1)
	}

	fmt.Println(string(j1))
	fmt.Println(string(j2))

}

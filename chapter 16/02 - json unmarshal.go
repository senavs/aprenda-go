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
	var j1 []byte = []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente Secreto","ContanBancaria":1000000.5}`)

	var p1 pessoa
	json.Unmarshal([]byte(j1), &p1)
	fmt.Println(p1)
	fmt.Println(p1.Nome)

}

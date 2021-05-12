package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type dentista struct {
	pessoa
	dentesArrancados int
	salario          float64
}

type arquiteto struct {
	pessoa
	tipoConstrucao string
	tamanhoLoucura string
}

type profissional interface {
	saudacao()
}

func (d dentista) saudacao() {
	fmt.Println("Olá, bom dia!", "Eu", d.nome, "já arraquei", d.dentesArrancados, "dentes")
}

func (a arquiteto) saudacao() {
	fmt.Println("Olá, bom dia!", "Eu", a.nome, "faço construições do tipo", a.tipoConstrucao)
}

func mandarSaudar(ps ...profissional) {
	for _, p := range ps {
		p.saudacao()
	}
}

func main() {
	mrDente := dentista{
		pessoa: pessoa{
			nome:  "Mr. Dente",
			idade: 30,
		},
		dentesArrancados: 987,
		salario:          8271.50,
	}

	mrPredio := arquiteto{
		pessoa: pessoa{
			nome:  "Mr. Prédio",
			idade: 50,
		},
		tipoConstrucao: "Verfical",
		tamanhoLoucura: "D+",
	}

	// métodos
	mrDente.saudacao()
	mrPredio.saudacao()

	fmt.Println()

	// interface
	mandarSaudar(mrDente, mrPredio)
}

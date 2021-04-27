package main

import (
	"fmt"
)

func main() {
	interpreterStringLiteral := "Olá, galerinha do Youtube!\nTudo bom com vocês??"
	rawStringLiteral := `Essa é uma string crua
	\n\n\n\t\t\t\t\t\t\n\n\n\n\n\n
	
	
	
	
	
	
	
	Consigo fazer uma zona aqui.
	`

	fmt.Println(interpreterStringLiteral)
	fmt.Println(rawStringLiteral)
}

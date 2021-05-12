package main

import (
	"fmt"
)

func main() {
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "queijos", "marg"}

	// slices from slices
	fatia := sabores[1:2]
	// fatia := sabores[:2]
	// fatia := sabores[1:]
	// fatia := sabores[:]

	fmt.Println(sabores)
	fmt.Println(fatia)

	// delete from slices
	// sabores = sabores[:3]
	sabores = append(sabores[:2], sabores[3:]...)
}

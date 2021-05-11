package main

const x = 10   // definido somente quando for usado
var y int = 10 // definido na declaração

var a int
var b float64

// várias constantes
const (
	c1 int     = 100
	c3 float64 = 100.40
)

func main() {
	// a = x  // x será int
	// b = x  // x será float

	// a = y  // sucesso
	// b = y  // erro
}

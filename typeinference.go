package main

import "fmt"

func main() {
	var i int
	var j = i // type int
	fmt.Printf("Tipo: %T Valor: %v", j, i)

	var x = 0.64 + 0.3i
	fmt.Printf("\nTipo: %T Valor: %v", x, x)
}

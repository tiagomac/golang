package main

import "fmt"

func main() {
	println(soma(3,2))
	fmt.Println(calcular(2))
	var foo (int, int) := calcular(2)
	fmt.Printf("Tipo: %T   -   Valor: %v", foo, foo)
}

func soma(x int, y int) int {
	return x + y;
}

func calcular(a int) (int, int){
	var quadrado int = a * a
	var cubo int = a * a * a
	return quadrado, cubo
}

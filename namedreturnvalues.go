package main

import "fmt"

func main() {
	fmt.Println(retornoNomeado(2))
}

func retornoNomeado(a int) (multitwo int, multithree int) {
	multitwo = a * 2;
	multithree = a * 3
	return
}


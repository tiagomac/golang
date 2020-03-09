package main

import (
	"fmt"
	"strings"
)

func main() {
	tabuleiro := [][]string {
		[]string{"X", "-", "-"},
		[]string{"O", "X", "O"},
		[]string{"-", "O", "X"}}

	fmt.Println(tabuleiro[0]) // first line
	fmt.Println(tabuleiro)
	for idx := 0; idx < len(tabuleiro); idx++ {
		fmt.Println(tabuleiro[idx])
	}

	fmt.Println("\n\n\n")

	for idx := 0; idx < len(tabuleiro); idx++ {
		fmt.Printf("%s", strings.Join(tabuleiro[idx], ","))
	}

}
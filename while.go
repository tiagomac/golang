package main

import "time"

func main() {

	var soma int = 2

	for ; soma < 600; {
		soma += soma
	}
	println(soma)

	for soma < 2000 { // loop while
		soma += soma
	}
	println(soma)

	for { // infite loop
		soma += soma
		println(soma)
		time.Sleep(100 * time.Millisecond)
	}

}
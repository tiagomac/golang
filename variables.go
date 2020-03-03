package main

import "fmt"

func main() {
	// types
	// int, int8, int 16, int 32, int64, uint, uint8, uint16, uint32 uint64, string
	// float64, float32 -- float64 is default
	// bool
	// complex64, complex128
	// byte, rune

	var age, height int = 20, 186
	var name string = "Tiago"
	println(age)
	println("My age is", age, "and my name is", name, "and my height is", height)

	var (
		surname string = "Mesquita"
		wallColor string = "white"
	)
	println("My surname is", surname, "and my wall color is", wallColor)

	var (
		test = "test"
		foo = 20
	)

	println(test, foo)

	const test2 = "HELLO CONST"
	println(test2)

	skyColor := "blue"
	println(skyColor)

	foo1,foo2,foo3 := 20,30,"foo3"
	println(foo1,foo2,foo3)

	var height2 = 1.34
	// %T for types and %v for value
	fmt.Printf("\nTipo: %T Valor: %v", height2, height2)



}

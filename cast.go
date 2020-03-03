package main

import "fmt"

func main(){
	var b float64 = 6.4
	//var c int = b --impossible

	var c int = int(b)
	var d float32 = float32(b)

	println(c)
	fmt.Printf("%v", d)


}

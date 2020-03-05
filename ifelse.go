package main

func main(){

	// a is undefined out of this scope
	if a := 11; a > 10 {
		println("a is bigger than 10")
	} else if (a > 5){
		println("A is bigger than 5")
	} else {
		println("a is lower than 5")
	}

}

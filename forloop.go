package main

func main() {
	var soma int = 0;

	for i := 1; i <= 10; i = i + 1 {
		soma = soma + i
	}

	println(soma)
}

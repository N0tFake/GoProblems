package main

import "fmt"

func fibonacci(n int) int {
	a := 0
	b := 1
	var aux int

	for i := 0; i < n; i++ {
		aux = b

		b = a + b
		a = aux
	}

	return a
}

func main() {
	var num int

	fmt.Print("Numero: ")
	fmt.Scanf("%d", &num)

	result := fibonacci(num)
	fmt.Printf("Resultado: %d\n", result)
}

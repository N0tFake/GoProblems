package main

/* Dado um numero, mostre o resultado do seu Fatorial */

import "fmt"

func main() {
	var num int

	fmt.Print("Numero: ")
	fmt.Scanf("%d", &num)

	result := 1
	for i := num; i > 0; i-- {
		result = result * i
	}

	fmt.Printf("O Fatorial de %d é: %d\n", num, result)

}

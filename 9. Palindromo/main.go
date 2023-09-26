package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindromo(s string) bool {
	str := strings.ReplaceAll(s, " ", "")
	tamanho := len(str)

	for i := 0; i < tamanho; i++ {
		if str[i] != str[tamanho-1-i] {
			return false
		}
	}
	return true
}
func main() {
	input := bufio.NewScanner(os.Stdin)

	fmt.Println("Palavra ou Frase:")
	input.Scan()
	phrase := input.Text()

	result := isPalindromo(phrase)

	if result {
		fmt.Println("A Palvra/Frase é um Palindromo!!")
	} else {
		fmt.Println("A Palvra/Frase NÃO é um Palindromo!!")
	}
}

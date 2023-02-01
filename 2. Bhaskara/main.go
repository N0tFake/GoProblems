package main

import (
	"errors"
	"fmt"
	"math"
)

// Dado a, b e c. Sendo eles representados por uma função de 2° grau(ax²+bx-c).
// Achar o valor de x' e x" usando a formula de bhaskara

func bhaskara(a, b, c float64) (x1 float64, x2 float64, erro error) {

	delta := ((math.Pow(b, 2)) - (4 * a * c))

	if delta < 0 || a == 0 {

		if delta < 0 {
			erro = errors.New("O delta é negativo. Equação não possui raíses reais.")
		} else if a == 0 {
			erro = errors.New("a = 0, logo 2*a = 0, sendo assim não existe divisão por 0.")
		}

		x1 = 0
		x2 = 0

		return
	}

	raiz := math.Sqrt(delta)

	x1 = (-b + raiz) / (2 * a)
	x2 = (-b - raiz) / (2 * a)

	return
}

func main() {
	x1, x2, erro := bhaskara(3, 3, -1)

	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Printf("X1: %.2f\nX2: %.2f\n", x1, x2)
	}

}

package main

import (
	"fmt"
	"math"
)

// Calcular a aréa das seguintes formas geometricas:
// Retangulo, triangulo, circulo, trapezio, losango

func areaRectangle(base, height float64) float64 {
	return base * height
}

func areaTriangle(base, height float64) float64 {
	return (base * height) / 2
}

func areaCircle(raio float64) float64 {
	return math.Pi * raio
}

func areaTrapeze(baseMaior, baseMenor, height float64) float64 {
	return ((baseMaior + baseMenor) * height) / 2
}

func areaLosango(diagonalMaior, diagonalMenor float64) float64 {
	return (diagonalMaior * diagonalMenor) / 2
}

func main() {
	rectangle := areaRectangle(10, 20)
	triangle := areaTriangle(10, 30)
	circle := areaCircle(3)
	trapeze := areaTrapeze(20, 10, 30)
	losango := areaLosango(20, 10)

	fmt.Printf("Rectangle: %.2f\nTriangle: %.2f\nCircle: %.2f\nTrapeze: %.2f\nLosango: %.2f", rectangle, triangle, circle, trapeze, losango)
}

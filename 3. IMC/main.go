package main

import (
	"fmt"
	"math"
)

// Dado um array ou slice de Pessoa, sendo que Pessoa é um struct.
// Calcular a media dos IMC do array/slice, encontrar a pessoa com o maior IMC e a co9.m o menor
// E mostrar a porcentagem de cada uma das classificações da tabela de imc encontradas no array/slice

// Tabela IMC
// menor que 18.5 - magreza
// entre 18.5 e 24.9 - Normal
// entre 25 e 29.9 - Sobrepeso
// entre 30 e 39.9 - Obesidade
// mairo que 40 - Obesidade Grave

type Pessoa struct {
	nome           string
	height         float64
	weight         float64
	imc            float64
	classification string
}

func imc(weight, height float64) (imc float64, classification string) {
	imc = (weight / math.Pow(height, 2))

	if imc < 18.5 {
		classification = "Magreza"
	} else if imc >= 18.5 && imc <= 24.9 {
		classification = "Normal"
	} else if imc >= 25 && imc <= 29.9 {
		classification = "Sobrepeso"
	} else if imc >= 30 && imc <= 39.9 {
		classification = "Obesidade"
	} else if imc >= 40 {
		classification = "Obesidade Grave"
	}

	return

}

func getMedia(pessoas *[]Pessoa) (media float64) {
	fmt.Println(pessoas)
	return
}

func calcImcs(pessoas []Pessoa) ([]Pessoa, Pessoa, Pessoa, float64) {

	var maior Pessoa
	var menor Pessoa

	for i, pessoa := range pessoas {
		pessoas[i].imc, pessoas[i].classification = imc(pessoa.weight, pessoa.height)

		if i == 0 {
			maior = pessoa
			menor = pessoa
		}

		if pessoa.imc > maior.imc {
			maior = pessoa
		}
		if pessoa.imc < menor.imc {
			menor = pessoa
		}
	}

	var soma float64 = 0.0

	for _, pessoa := range pessoas {
		soma = pessoa.imc
	}

	media := soma / float64(len(pessoas))

	return pessoas, maior, menor, media

}

func getPorcGeral(pessoas []Pessoa) (porcMagreza, porcNormal, porcSobrepeso, porcObesidade, porcObesidadeG float64) {

	var (
		countMagreza    int = 0
		countNormal     int = 0
		countSobrepeso  int = 0
		countObesidade  int = 0
		countObesidadeG int = 0
	)

	for _, pessoa := range pessoas {
		if pessoa.classification == "Magreza" {
			countMagreza++
		} else if pessoa.classification == "Normal" {
			countNormal++
		} else if pessoa.classification == "Sobrepeso" {
			countSobrepeso++
		} else if pessoa.classification == "Obesidade" {
			countObesidade++
		} else if pessoa.classification == "Obesidade Grave" {
			countObesidadeG++
		}
	}

	length := len(pessoas)
	porcMagreza = getPorc(countMagreza, length)
	porcNormal = getPorc(countNormal, length)
	porcSobrepeso = getPorc(countSobrepeso, length)
	porcObesidade = getPorc(countObesidade, length)
	porcObesidadeG = getPorc(countObesidadeG, length)

	return
}

func getPorc(count, length int) float64 {
	if count != 0 {
		return float64((count * 100) / length)
	}
	return 0.0
}

func main() {
	p1 := Pessoa{nome: "Fulano", height: 1.76, weight: 67.4, imc: 0.0, classification: ""}
	p2 := Pessoa{nome: "Fulano", height: 1.53, weight: 87.4, imc: 0.0, classification: ""}
	p3 := Pessoa{nome: "Fulano", height: 1.81, weight: 39.2, imc: 0.0, classification: ""}
	p4 := Pessoa{nome: "Fulano", height: 1.68, weight: 71.3, imc: 0.0, classification: ""}

	pessoas := []Pessoa{p1, p2, p3, p4}

	pessoas, maior, menor, media := calcImcs(pessoas)
	porcMagreza, porcNormal, porcSobrepeso, porcObesidade, porcObesidadeG := getPorcGeral(pessoas)

	fmt.Println("Pessoa com o maior IMC: ", maior)
	fmt.Println("Pessoa com o menor IMC: ", menor)
	fmt.Println("Media dos IMCs: ", media)
	fmt.Println("---------------- Porcentagens ----------------")
	fmt.Printf("Magreza: %.2f\n", porcMagreza)
	fmt.Printf("Normal: %.2f\n", porcNormal)
	fmt.Printf("Sobrepeso: %.2f\n", porcSobrepeso)
	fmt.Printf("Obesidade: %.2f\n", porcObesidade)
	fmt.Printf("Obesiadede Grave: %.2f\n", porcObesidadeG)

}

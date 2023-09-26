package main

import "fmt"

/*
Dada uma matriz de números inteiros 'nums' e um número inteiro 'target', retorne os índices dos dois números de forma que a soma deles sejatarget .
Você pode assumir que cada entrada teria exatamente uma solução e não pode usar o mesmo elemento duas vezes.
Você pode retornar a resposta em qualquer ordem.
*/

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, ok := numMap[complement]; ok {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	// nums := []int{3, 2, 4}
	// nums := []int{3, 3}
	// nums := []int{9, 11, 5, 8}

	target := 13
	result := twoSum(nums, target)

	fmt.Println(result)
}

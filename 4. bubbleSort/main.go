package main

import (
	"fmt"
)

func bubbleSort(arr []int) []int {

	for i := range arr {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				aux := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = aux
			}
		}
	}

	return arr
}

func main() {
	arr := []int{2, 3, 5, 7, -1, 2, 12, 1, 2, -353, 1, 992}
	result := bubbleSort(arr)

	fmt.Println(result)
}

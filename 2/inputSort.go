package main

import "fmt"

func main() {
	arr := []int{5, 3, 6, 1, 2}
	inputSort(arr)
	fmt.Println(arr)
}

func inputSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}
}

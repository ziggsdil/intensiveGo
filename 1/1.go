package main

import "fmt"

func main() {
	arr := []int{5, 3, 6, 1}

	bubbleSort(arr)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
}

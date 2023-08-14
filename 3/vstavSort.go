package main

import "fmt"

func main() {
	arr := []int{5, 3, 2, 6, 1, 23}
	insertSort(arr)
	fmt.Println(arr)
}

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

package main

import "fmt"

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(findMaxSubArr(arr))
}

// максимальную сумму подмассива
func findMaxSubArr(arr []int) int {
	sum := 0
	currentSum := 0
	for i := 0; i < len(arr); i++ {
		currentSum += arr[i]
		sum = max(currentSum, sum)

		if currentSum < 0 {
			currentSum = 0
		}
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

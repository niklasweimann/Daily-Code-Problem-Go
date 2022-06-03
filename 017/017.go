package main

import (
	"fmt"
)

func main() {
	list := []int{-9, -2, 0, 2, 3}
	fmt.Println(squareAndOrder(list))
}

func squareAndOrder(arr []int) []int {
	n := len(arr)
	left, right := 0, n-1
	index := n - 1
	result := make([]int, len(arr))

	for index >= 0 {
		if abs(arr[left]) >= abs(arr[right]) {
			result[index] = arr[left] * arr[left]
			left += 1
		} else {
			result[index] = arr[right] * arr[right]
			right -= 1
		}
		index -= 1
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

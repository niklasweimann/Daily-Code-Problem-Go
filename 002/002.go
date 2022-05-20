package main

import "fmt"

func main() {
	var list = []int{1, 2, 3, 4, 5}
	fmt.Println(findSum(list, 9)) // 2 3 4

	list = []int{1, 2, 3, 4, 5}
	fmt.Println(findSum(list, 5)) // 2 3

	list = []int{4, 5, 7, 8, 112}
	fmt.Println(findSum(list, 127)) // 7 8 112
}

func findSum(list []int, k int) []int {
	start := 0
	for i := 0; i < len(list); i++ {
		sum := 0
		for j := start; j < k; j++ {
			sum += list[j]
			if sum > k {
				break
			}
			if sum == k {
				return list[start : j+1]
			}
		}
		start += 1
	}
	return []int{}
}

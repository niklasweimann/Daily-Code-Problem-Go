package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(rotate(list, 2))
}

func rotate(list []int, shiftSize int) []int {
	for i := 0; i < shiftSize; i++ {
		temp := list[0]
		for j := 0; j < len(list)-1; j++ {
			list[j] = list[j+1]
		}
		list[len(list)-1] = temp
	}
	return list
}

package main

import (
	"fmt"
	"math"
)

func main() {
	list := [][]int{{0, 3}, {2, 6}, {3, 4}, {6, 9}}
	fmt.Println(findSmallestSet(list)) // expected {3, 6}
}

func findSmallestSet(list [][]int) (int, int) {
	smallMax := math.MaxInt
	largeMin := -math.MaxInt
	for _, v := range list {
		if v[1] < smallMax {
			smallMax = v[1]
		}
		if v[0] > largeMin {
			largeMin = v[0]
		}
	}
	return smallMax, largeMin
}

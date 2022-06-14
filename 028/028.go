package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getSquareRoot(9))

}

func getSquareRoot(i int) int {
	if i == 0 || i == 1 {
		return i
	}

	start, end := 0, math.MaxInt
	for start+1 < end {
		mid := start + (end-start)/2
		if mid == i/mid {
			return mid
		}
		if mid > i/mid {
			end = mid
		} else {
			start = mid
		}
	}
	return start
}

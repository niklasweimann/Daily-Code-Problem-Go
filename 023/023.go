package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getNumberOfCounts(10))
}

func getNumberOfCounts(n int) int {
	return int(math.Ceil(math.Log2(float64(n)) + 1))
}

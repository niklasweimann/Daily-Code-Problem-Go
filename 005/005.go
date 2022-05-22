package main

import "fmt"

func main() {
	steps := []int{2, 0, 1, 0}
	fmt.Println(canReachEnd(0, steps))
	steps = []int{1, 1, 0, 1}
	fmt.Println(canReachEnd(0, steps))
}

func canReachEnd(current int, steps []int) bool {
	// Stop at the end
	if current == len(steps)-1 {
		return true
	}
	// if we are on a 0 position, we are stuck
	if steps[current] == 0 {
		return false
	}

	// iterate list and try to reach the end using recursion
	rv := false
	for jump := 1; jump <= steps[current]; jump++ {
		if current+jump < len(steps) {
			rv = rv || canReachEnd(current+jump, steps)
		}
	}
	return rv
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	a_1 := "abcde"
	b_1 := "cdeab"
	fmt.Println(isShiftEqual(a_1, b_1)) // expected true
	a_2 := "abc"
	b_2 := "acb"
	fmt.Println(isShiftEqual(a_2, b_2)) // expected false
}

func isShiftEqual(a string, b string) bool {
	return strings.Contains(a+a, b)
}

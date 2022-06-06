package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPalindrome("waterrfetawx", 2))
}

func isPalindrome(s string, i int) bool {
	return (isKPalRec(s, reverse(s), len(s), len(s)) <= i*2)
}

func isKPalRec(s1, s2 string, len1 int, len2 int) int {
	if len1 == 0 {
		return len2
	}

	if len2 == 0 {
		return len1
	}

	if s1[len1-1] == s2[len2-1] {
		return isKPalRec(s1, s2, len1-1, len2-1)
	}

	return 1 + int(math.Min(float64(isKPalRec(s1, s2, len1-1, len2)), float64(isKPalRec(s1, s2, len1, len2-1))))
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

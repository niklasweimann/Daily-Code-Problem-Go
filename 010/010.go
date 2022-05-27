package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	W := "ab"
	S := "abxaba"
	fmt.Println(anagramIndices(W, S))
}

func is_anagram(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	str1 = sort_string(str1)
	str2 = sort_string(str2)

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false

		}
	}
	return true
}

func sort_string(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func anagramIndices(word string, s string) []int {
	result := []int{}
	for i := 0; i < len(s)-len(word)+1; i += 1 {
		window := s[i : i+len(word)]
		if is_anagram(window, word) {
			result = append(result, i)
		}
	}
	return result
}

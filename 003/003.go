package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Test 1: ")
	chars := []string{"a", "e", "i"}
	list := "figehaeci"
	result := *findSubstring(chars, list)
	fmt.Println(result)

	fmt.Print("Test 2: ")
	chars2 := []string{"l", "o", "h"}
	list2 := "hallohallohallo"
	result2 := *findSubstring(chars2, list2)
	fmt.Println(result2)
}

func findSubstring(chars []string, list string) *string {
	startIndex := strings.Index(list, chars[0])
	// we will not anything if the list does not contain the chars or
	// the lis is too short
	if startIndex == -1 || len(list)-startIndex < len(chars) {
		return nil
	}
	return parseString(chars, list[startIndex:])
}

func parseString(chars []string, substring string) *string {
	charIndex := 0
	bestValue := ""
	current := ""
	for i := 0; i < len(substring); i++ {
		// start check of subset if we find a better solution
		if string(substring[i]) == chars[0] && charIndex > 0 && len(substring) > i+1 {
			localResult := parseString(chars, substring[i:])
			// check if we found a better solution
			if localResult != nil && len(*localResult) < len(bestValue) && len(*localResult) >= len(chars) {
				bestValue = *localResult
			}
		}

		// Build current string
		if charIndex < len(chars) {
			current += string(substring[i])
			if string(substring[i]) == chars[charIndex] {
				charIndex = charIndex + 1
			}
		}

		// Check for better solution
		if charIndex == len(chars) {
			// Always set bestValue if this is the first bestValue
			if len(bestValue) == 0 {
				bestValue = current
			}
			// Only update bastValue if there is a shorter solution
			if len(bestValue) > 1 && len(current) < len(bestValue) {
				bestValue = current
			}

		}
	}

	// return nil if no solution was found
	if len(bestValue) > 0 {
		return &bestValue
	} else {
		return nil
	}
}

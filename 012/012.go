package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseString("hello world here")) // expected "here world hello"
}

func reverseString(text string) string {
	splitted := strings.Split(text, " ")
	for i, j := 0, len(splitted)-1; i < j; i, j = i+1, j-1 {
		splitted[i], splitted[j] = splitted[j], splitted[i]
	}
	return strings.Join(splitted, " ")
}

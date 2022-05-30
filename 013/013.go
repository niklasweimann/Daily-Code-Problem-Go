package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(IsLetter(rune(":"[0])))
	fmt.Println(reverseString("hello/world:here")) // expected "here/world:hello"
	fmt.Println(reverseString("hello/world:here/"))
	fmt.Println(reverseString("hello//world:here"))
}

func reverseString(text string) string {
	words := []string{}
	curWord := ""
	for i := 0; i < len(text); i++ {
		if !IsLetter(rune(text[i])) {
			words = append(words, curWord)
			curWord = ""
			continue
		}
		curWord += string(text[i])
	}
	words = append(words, curWord)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	result := ""
	indexWord := 0
	i := 0
	for i < len(text) {
		if IsLetter(rune(text[i])) {
			result += words[indexWord]
			i += len(words[indexWord]) - 1
			if indexWord < len(words)-1 {
				indexWord++
			}
			continue
		}
		if !IsLetter(rune(text[i])) {
			result += string(text[i])
			i += 1
		}
	}
	return result
}

var IsLetter = unicode.IsLetter

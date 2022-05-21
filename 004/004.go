package main

import (
	"fmt"
)

func main() {
	fmt.Print("Test 1: ")
	link := List{}
	link.Insert("a").Insert("b").Insert("b").Insert("a")
	result := isListPalindrom(link)
	fmt.Println(result)

	fmt.Print("Test 1: ")
	link = List{}
	link.Insert("a").Insert("b").Insert("a").Insert("a")
	result = isListPalindrom(link)
	fmt.Println(result)
}

func isListPalindrom(list List) bool {
	s := NewStack()
	current := list.head
	for ok := true; ok; ok = current != nil {
		s.Push(current.key)
		current = current.next
	}

	current = list.head
	for ok := true; ok; ok = current != nil {
		x := s.Pop()
		if x != current.key {
			return false
		}
		current = current.next
	}
	return true
}

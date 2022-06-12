package main

import (
	"fmt"

	singlelinkedlist "github.com/niklasweimann/Daily-Code-Problem-Go/basic/single-linked-list"
)

func main() {
	list99 := singlelinkedlist.NewList[int]().Insert(9).Insert(9)
	list52 := singlelinkedlist.NewList[int]().Insert(5).Insert(2)

	fmt.Println(addLists(list99, list52))
}

func addLists[T int](list1, list2 *singlelinkedlist.List[int]) *singlelinkedlist.List[int] {
	if list1.Length != list2.Length {
		panic("Lists need to be of equal length")
	}
	current1 := list1.Head
	current2 := list2.Head
	result := &singlelinkedlist.List[int]{}
	carry := 0
	for current1 != nil || current2 != nil || carry > 0 {
		value1, value2 := 0, 0
		if current1 != nil {
			value1 = current1.Value
			current1 = current1.Next
		}
		if current2 != nil {
			value2 = current2.Value
			current2 = current2.Next
		}
		sum := carry + value1 + value2
		carry = sum / 10
		result.Insert(sum % 10)
	}

	return result
}

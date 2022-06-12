package singlelinkedlist

import (
	"fmt"
	"strings"
)

type (
	List[T any] struct {
		Length int
		Head   *listNode[T]
	}
	listNode[T any] struct {
		Value T
		Next  *listNode[T]
	}
)

func NewList[T any]() *List[T] {
	return &List[T]{0, nil}
}

func (list *List[T]) Insert(value T) *List[T] {
	list.Length += 1
	if list.Head == nil {
		list.Head = &listNode[T]{value, nil}
		return list
	}
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &listNode[T]{value, nil}
	return list
}

func (list *List[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[ ")
	current := list.Head
	if current != nil {
		sb.WriteString(fmt.Sprintf("%d ", current.Value))
	}
	for current.Next != nil {
		current = current.Next
		sb.WriteString(fmt.Sprintf("%d ", current.Value))
	}
	sb.WriteString(" ]")
	return sb.String()
}

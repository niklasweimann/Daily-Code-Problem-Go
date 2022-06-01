package main

import (
	"fmt"
)

func main() {
	fmt.Println(generate[int]().Left().Left().Right().Value)
}

func generate[T any]() Node[T] {
	return Node[T]{}
}

// Node definition
type (
	Node[T any] struct {
		left  *Node[T]
		right *Node[T]
		Value T

		left_created  bool
		right_created bool
	}
)

func (n Node[T]) Left() *Node[T] {
	if !n.left_created {
		n.left = &Node[T]{}
	}
	n.left_created = true
	return n.left
}

func (n Node[T]) Right() *Node[T] {
	if !n.right_created {
		n.right = &Node[T]{}
	}
	n.right_created = true
	return n.right
}

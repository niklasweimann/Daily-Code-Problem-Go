package main

import (
	"fmt"
	"github.com/niklasweimann/Daily-Code-Problem-Go/basic/queue"
	"github.com/niklasweimann/Daily-Code-Problem-Go/basic/tree"
)

func main() {
	fiveNode := tree.Node[int]{
		Value: 5,
	}
	fourNode := tree.Node[int]{
		Value: 4,
	}
	threeNode := tree.Node[int]{
		Value: 3,
		Left:  &fourNode,
		Right: &fiveNode,
	}
	twoNode := tree.Node[int]{
		Value: 2,
	}
	r := tree.Node[int]{
		Parent: nil,
		Left:   &twoNode,
		Right:  &threeNode,
		Value:  1,
	}
	BFS[int](r)
}

func BFS[T any](root tree.Node[T]) {
	q := queue.Queue[tree.Node[T]]{}
	q.Enqueue(root) //You don't need to write the root here, it will be written in the loop
	for q.Length > 0 {
		n := q.Dequeue()
		fmt.Print(n.Value)
		fmt.Print(", ")
		if n.Left != nil {
			q.Enqueue(*n.Left) //enqueue the left child
		}
		if n.Right != nil {
			q.Enqueue(*n.Right) //enque the right child
		}
	}
}

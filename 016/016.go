package main

import (
	"fmt"

	"github.com/niklasweimann/Daily-Code-Problem-Go/basic/tree"
)

func main() {
	fiveNode := tree.Node[int]{
		Value: 1,
	}
	fourNode := tree.Node[int]{
		Value: 1,
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
		Value:  1234,
	}
	fmt.Println(minLevelSum(&r)) // should return 3
}

func minLevelSum(root *tree.Node[int]) int {
	minimum, minLevel, level := root.Value, 1, 0
	queue := []*tree.Node[int]{root}

	for len(queue) != 0 {
		sum, n := 0, len(queue)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Value

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		level++
		if sum < minimum {
			minimum, minLevel = sum, level
		}
	}
	return minLevel
}

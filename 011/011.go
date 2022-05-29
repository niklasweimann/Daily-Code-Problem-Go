package main

import (
	"fmt"

	"github.com/niklasweimann/Daily-Code-Problem-Go/basic/tree"
)

func main() {
	right2 := tree.Node[int]{
		Value: 5,
	}
	left2 := tree.Node[int]{
		Value: 4,
	}
	right1 := tree.Node[int]{
		Value: 3,
		Left:  &left2,
		Right: &right2,
	}
	left1 := tree.Node[int]{
		Value: 2,
	}
	root := tree.Node[int]{
		Value: 1,
		Left:  &left1,
		Right: &right1,
	}

	fmt.Println(parseTree(&root, 4, 5)) //expected 3
	fmt.Println(parseTree(&root, 2, 5)) //expected 1
}

func parseTree(root *tree.Node[int], value1 int, value2 int) int {
	value1path := []int{}
	value2path := []int{}

	if !findPath(root, &value1path, value1) || !findPath(root, &value2path, value2) {
		return -1
	}

	i := 0
	for i < len(value1path) && i < len(value2path) {
		if value1path[i] != value2path[i] {
			break
		}
		i += 1
	}
	return value1path[i-1]
}

func findPath(root *tree.Node[int], path *[]int, k int) bool {
	if root == nil {
		return false
	}

	*path = append(*path, root.Value)

	if root.Value == k {
		return true
	}

	if (root.Left != nil && findPath(root.Left, path, k)) || (root.Right != nil && findPath(root.Right, path, k)) {
		return true
	}

	*path = (*path)[:len(*path)-1]
	return false
}

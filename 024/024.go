package main

import (
	"fmt"

	"github.com/niklasweimann/Daily-Code-Problem-Go/basic/tree"
)

func main() {
	root := &tree.Node[int32]{Value: 10}
	root = tree.InsertNode(root, 5)
	root = tree.InsertNode(root, 15)
	root = tree.InsertNode(root, 15)
	root = tree.InsertNode(root, 11)
	isPairPresent(root, 20) // Expected: 5 + 15 = 20

}
func treeToList(node *tree.Node[int32], list *[]int) *[]int {
	// Base Case
	if node == nil {
		return list
	}
	treeToList(node.Left, list)
	*list = append(*list, int(node.Value))
	treeToList(node.Right, list)
	return list
}

func isPairPresent(node *tree.Node[int32], target int) {
	a1 := []int{}
	a2 := *treeToList(node, &a1)
	start := 0
	end := len(a2) - 1

	for start < end {

		if a2[start]+a2[end] == target {
			fmt.Printf("Found: %d + %d = %d", a2[start], a2[end], target)
			return
		}

		if a2[start]+a2[end] > target {
			end--
		}

		if a2[start]+a2[end] < target {
			start++
		}
	}

	fmt.Println("No pair found!")
}

package main

import (
	"fmt"
	"strconv"
	"strings"

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
	printPath(&root, "")
}

func printPath(node *tree.Node[int], path string) {
	if node == nil {
		return
	} else if node.Left == nil && node.Right == nil {
		fmt.Println(path, node.Value)
	} else {
		printPath(node.Left, strings.Join([]string{path, strconv.Itoa(node.Value)}, " "))
		printPath(node.Right, strings.Join([]string{path, strconv.Itoa(node.Value)}, " "))
	}
}

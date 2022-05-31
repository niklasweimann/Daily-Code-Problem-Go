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
	fmt.Println(isSubtree(&root, &right1))
}

func isSubtree(s *tree.Node[int], t *tree.Node[int]) bool {
	if s == nil {
		return false
	}

	if s.Value == t.Value {
		if isSame(s, t) {
			return true
		}
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(root1 *tree.Node[int], root2 *tree.Node[int]) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Value != root2.Value {
		return false
	}

	return isSame(root1.Left, root2.Left) && isSame(root1.Right, root2.Right)
}

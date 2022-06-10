package tree

type (
	Node[T any] struct {
		Parent *Node[T]
		Left   *Node[T]
		Right  *Node[T]
		Value  T
	}
)

func InsertNode(root *Node[int32], value int32) *Node[int32] {
	if root == nil {
		return &Node[int32]{Value: value, Left: nil, Right: nil}
	}
	if value < root.Value {
		root.Left = InsertNode(root.Left, value)
	} else {
		root.Right = InsertNode(root.Right, value)
	}
	return root
}

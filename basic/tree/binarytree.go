package tree

type (
	Node[T any] struct {
		Parent *Node[T]
		Left   *Node[T]
		Right  *Node[T]
		Value  T
	}
)

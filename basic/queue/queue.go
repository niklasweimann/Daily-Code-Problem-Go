package queue

type (
	Queue[T any] struct {
		head   *Node[T]
		Length int
	}
	Node[T any] struct {
		value T
		next  *Node[T]
	}
)

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{nil, 0}
}

func (Q *Queue[T]) Enqueue(value T) *Queue[T] {
	newNode := new(Node[T])
	newNode.value = value
	Q.Length += 1
	if Q.head == nil {
		Q.head = newNode
		return Q
	}

	c := Q.head
	for c.next != nil {
		c = c.next
	}

	c.next = newNode
	return Q
}

func (Q *Queue[T]) Dequeue() T {
	if Q.head == nil {
		return *new(T)
	}

	Q.Length -= 1
	c := Q.head
	Q.head = Q.head.next
	return c.value
}

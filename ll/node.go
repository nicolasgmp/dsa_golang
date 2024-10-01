package ll

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{val, nil}
}

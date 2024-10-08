package ll

import "fmt"

type SinglyLinkedList[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

func NewLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{Head: nil, Tail: nil, Length: 0}
}

func (ll *SinglyLinkedList[T]) Push(val T) {
	newNode := NewNode(val)
	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}
	ll.Length++
}

func (ll *SinglyLinkedList[T]) Pop() T {
	if ll.Length == 0 {
		fmt.Println("Empty Linked List")
		var zero T
		return zero
	}

	if ll.Head.Next == nil {
		curr := ll.Head
		ll.Head = nil
		ll.Tail = nil
		ll.Length--
		return curr.Val
	}

	curr := ll.Head
	pre := ll.Head

	for curr.Next != nil {
		pre = curr
		curr = curr.Next
	}

	ll.Tail = pre
	ll.Tail.Next = nil
	ll.Length--
	return curr.Val
}

func (ll *SinglyLinkedList[T]) Shift() T {
	if ll.Length == 0 {
		fmt.Println("Empty list")
		var zero T
		return zero
	}

	curr := ll.Head
	ll.Head = ll.Head.Next
	ll.Length--

	if ll.Length == 0 {
		ll.Tail = nil
	}

	return curr.Val
}

func (ll *SinglyLinkedList[T]) Unshift(val T) {
	newNode := NewNode(val)

	if ll.Length == 0 {
		ll.Head = newNode
		ll.Tail = newNode
	}

	newNode.Next = ll.Head
	ll.Head = newNode

	ll.Length++

	if ll.Length == 1 {
		ll.Tail = newNode
	}
}

func (ll *SinglyLinkedList[T]) Get(idx int) *Node[T] {
	if idx < 0 {
		fmt.Println("Index less than 0")
		var zero *Node[T]
		return zero
	}

	if idx >= ll.Length {
		fmt.Println("Index greater than LinkedList length")
		var zero *Node[T]
		return zero
	}

	i := 0
	curr := ll.Head
	for i < idx {
		curr = curr.Next
		i++
	}

	return curr
}

func (ll *SinglyLinkedList[T]) Set(idx int, val T) {
	if idx < 0 {
		fmt.Println("Index less than 0")
		return
	}

	if idx >= ll.Length {
		fmt.Println("Index greater than LinkedList length")
		return
	}

	curr := ll.Get(idx)
	curr.Val = val
}

func (ll *SinglyLinkedList[T]) Insert(idx int, val T) {
	newNode := NewNode(val)

	if idx < 0 {
		fmt.Println("Index less than 0")
		return
	}

	if idx >= ll.Length {
		fmt.Println("Index greater than LinkedList length")
		return
	}

	if idx == 0 {
		ll.Unshift(val)
	}

	pre := ll.Get(idx - 1)
	newNode.Next = pre.Next
	pre.Next = newNode

	ll.Length++
}

func (ll *SinglyLinkedList[T]) Remove(idx int) {

	if idx < 0 {
		fmt.Println("Index less than 0")
		return
	}

	if idx >= ll.Length {
		fmt.Println("Index greater than LinkedList length")
		return
	}

	if idx == 0 {
		ll.Shift()
	}

	pre := ll.Get(idx - 1)
	toRm := ll.Get(idx)
	pre.Next = toRm.Next

	ll.Length--
}

func (ll *SinglyLinkedList[T]) Reverse() {
	curr := ll.Head
	ll.Head = ll.Tail
	ll.Tail = curr
	var pre *Node[T]
	var next *Node[T]
	for i := 0; i < ll.Length; i++ {
		next = curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
}

package main

import (
	"dsa_estudo/ll"
	"fmt"
)

func main() {
	ll := ll.NewLinkedList[string]()
	ll.Push("hi")
	ll.Push("corinthians")
	ll.Push("garro")
	ll.Push("coronado")
	ll.Unshift("memphis")
	ll.Unshift("cássio")

	fmt.Println(ll.Pop())
	fmt.Println(ll.Pop())
	fmt.Println(ll.Shift())

	fmt.Println("Lista atual")
	curr := ll.Head
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

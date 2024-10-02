package main

import (
	"dsa_estudo/ll"
	"fmt"
)

func main() {
	ll := ll.NewLinkedList[string]()
	ll.Unshift("memphis")
	ll.Unshift("cássio")
	ll.Push("hi")
	ll.Push("corinthians")
	ll.Push("garro")
	ll.Push("coronado")
	ll.Set(2, "romero")

	fmt.Println("Lista Atual")
	curr := ll.Head
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

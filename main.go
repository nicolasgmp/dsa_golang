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

	fmt.Println(ll.Pop())
}

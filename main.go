package main

import (
	"fmt"

	"github.com/rianby64/data-structures-self-study/doublylinkedlist"
)

func main() {

	dll := doublylinkedlist.New()

	dll.Insert(33).Insert(44).Insert(55).Insert(66)

	ss := dll.Filter(func(d doublylinkedlist.DoublyLinkedList, i int) bool {
		fmt.Println(d, i)
		return true
	})
	// for curr := dll.Next(); curr != nil; curr = curr.Next() {
	// 	fmt.Println(curr)
	// }

	fmt.Println("finito", ss)
}

package main

import (
	"fmt"

	"github.com/rianby64/data-structures-self-study/list"
)

func main() {

	dll := list.New()

	dll.Insert(33).Insert(44).Insert(55).Insert(66)

	ss := dll.Filter(func(d list.List, i int) bool {
		fmt.Println(d, i)
		return true
	})

	found := dll.Find(func(d list.List, i int) bool {
		v := d.Value().(int)
		return v == 44
	})

	fmt.Println("finito", dll, ss, found)
}

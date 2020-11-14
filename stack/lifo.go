package lifo

import (
	"github.com/rianby64/data-structures-self-study/list"
)

// Stack interface
type Stack interface {
	Pop() interface{}
	Push(c interface{})
}

type lifo struct {
	list list.List
}

func (l *lifo) Pop() interface{} {
	last := l.list.Last()
	defer last.Delete()
	return last.Value()
}

func (l *lifo) Push(c interface{}) {
	item := l.list.Last()
	item.Insert(c)
}

// New constructs an stack on top of a lifo
func New() Stack {
	return &lifo{
		list: list.New(),
	}
}

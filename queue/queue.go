package queue

import (
	"github.com/rianby64/data-structures-self-study/list"
)

// Queue interface.
type Queue interface {
	Enqueue(v interface{})
	Dequeue() interface{}
}

type queue struct {
	list list.List
}

func (q *queue) Enqueue(v interface{}) {
	q.list.Insert(v)
}

func (q *queue) Dequeue() interface{} {
	last := q.list.Last()

	defer last.Delete()

	return last.Value()
}

// New constructor.
func New() Queue {
	return &queue{
		list: list.New(),
	}
}

package list

import (
	"github.com/rianby64/data-structures-self-study/cell"
)

type edges struct {
	first  *list
	last   *list
	length int
}

type list struct {
	payload cell.Cell
	edges   *edges
	next    *list
	prev    *list
}

type predicate func(item List, index int) bool

// List implements a list using a doubly-linked-list.
type List interface {
	cell.Cell
	Next() List
	First() List
	Last() List
	Insert(payload interface{}) List
	Delete() List
	Update(payload interface{}) List
	Length() int

	Filter(p predicate) List
	Find(p predicate) List
}

func isEmptyList(l *list) bool {
	return l.edges.first.Next() == nil
}

func isFirstInList(l *list) bool {
	return l.edges.first == l
}

func (l *list) SetValue(v interface{}) {
	l.payload.SetValue(v)
}

// Filter should be inside of an abstraction as this method doesn't belong to Linked-List.
func (l *list) Filter(p predicate) List {
	ll := New()
	i := 0

	for curr := l.edges.first.next; curr != nil && i < l.edges.length; curr = curr.next {
		if p(curr, i) {
			ll.Last().Insert(curr.payload.Value())
		}

		i++
	}

	return ll
}

// Find should be inside of an abstraction as this method doesn't belong to Linked-List.
func (l *list) Find(p predicate) List {
	i := 0

	for curr := l.edges.first.next; curr != nil && i < l.edges.length; curr = curr.next {
		if p(curr, i) {
			return curr
		}

		i++
	}

	return nil
}

func (l *list) Length() int {
	return l.edges.length
}

func (l *list) Value() interface{} {
	if l.payload != nil {
		return l.payload.Value()
	}

	return nil
}

func (l *list) Next() List {
	if l.next == nil {
		return nil
	}

	return l.next
}

func (l *list) First() List {
	if isEmptyList(l) {
		return l.edges.first
	}

	return l.edges.first.Next()
}

func (l *list) Last() List {
	return l.edges.last
}

func (l *list) Insert(payload interface{}) List {
	inserted := &list{
		edges:   l.edges,
		payload: cell.New(payload),
		next:    l.next,
		prev:    l,
	}

	if l.next == nil {
		l.edges.last = inserted
	} else {
		l.next.prev = inserted
	}

	l.next = inserted
	l.edges.length++

	return l.next
}

// perdio sentido esta funcion.
func (l *list) Update(payload interface{}) List {
	if isEmptyList(l) {
		return l.Insert(payload)
	}

	if isFirstInList(l) {
		l.next.payload.SetValue(payload)

		return l.next
	}

	l.payload.SetValue(payload)

	return l
}

func (l *list) Delete() List {
	l.edges.length--
	if l.edges.length < 0 {
		l.edges.length = 0
	}

	if l.prev != nil {
		if l.next != nil {
			l.next.prev = l.prev
			l.prev.next = l.next
		} else {
			l.prev.next = nil
			l.edges.last = l.prev
		}

		return l.prev
	}

	if l.next != nil {
		if l.next.next != nil {
			l.next.prev = l
		} else {
			l.edges.last = l
		}

		l.next = l.next.next
	}

	return l
}

// New is the constructor.
func New() List {
	l := list{}
	l.edges = &edges{
		first: &l,
		last:  &l,
	}

	return &l
}

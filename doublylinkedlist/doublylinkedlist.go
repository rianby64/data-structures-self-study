package doublylinkedlist

import "github.com/rianby64/data-structures-self-study/cell"

type edges struct {
	first  *doublylinkedlist
	last   *doublylinkedlist
	length int
}

type doublylinkedlist struct {
	payload cell.Cell
	edges   *edges
	next    *doublylinkedlist
	prev    *doublylinkedlist
}

type predicate func(item DoublyLinkedList, index int) bool

// DoublyLinkedList whatever
type DoublyLinkedList interface {
	cell.Cell
	Next() DoublyLinkedList
	First() DoublyLinkedList
	Last() DoublyLinkedList
	Insert(payload interface{}) DoublyLinkedList
	Delete() DoublyLinkedList
	Update(payload interface{}) DoublyLinkedList
	Length() int

	Filter(p predicate) DoublyLinkedList
	Find(p predicate) DoublyLinkedList
}

func isEmptyList(l *doublylinkedlist) bool {
	return l.edges.first.Next() == nil
}

func isFirstInList(l *doublylinkedlist) bool {
	return l.edges.first == l
}

func (l *doublylinkedlist) SetValue(v interface{}) {
	l.payload.SetValue(v)
}

// Filter should be inside of an abstraction as this method doesn't belong to Linked-List
func (l *doublylinkedlist) Filter(p predicate) DoublyLinkedList {
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

// Find should be inside of an abstraction as this method doesn't belong to Linked-List
func (l *doublylinkedlist) Find(p predicate) DoublyLinkedList {
	i := 0
	for curr := l.edges.first.next; curr != nil && i < l.edges.length; curr = curr.next {
		if p(curr, i) {
			return curr
		}
		i++
	}
	return nil
}

func (l *doublylinkedlist) Length() int {
	return l.edges.length
}

func (l *doublylinkedlist) Value() interface{} {
	if l.payload != nil {
		return l.payload.Value()
	}
	return nil
}

func (l *doublylinkedlist) Next() DoublyLinkedList {
	if l.next == nil {
		return nil
	}
	return l.next
}

func (l *doublylinkedlist) First() DoublyLinkedList {
	if isEmptyList(l) {
		return l.edges.first
	}
	return l.edges.first.Next()
}

func (l *doublylinkedlist) Last() DoublyLinkedList {
	return l.edges.last
}

func (l *doublylinkedlist) Insert(payload interface{}) DoublyLinkedList {
	inserted := &doublylinkedlist{
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

// perdio sentido esta funcion
func (l *doublylinkedlist) Update(payload interface{}) DoublyLinkedList {
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

func (l *doublylinkedlist) Delete() DoublyLinkedList {
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

// New is the constructor
func New() DoublyLinkedList {
	l := doublylinkedlist{}
	l.edges = &edges{
		first: &l,
		last:  &l,
	}
	return &l
}

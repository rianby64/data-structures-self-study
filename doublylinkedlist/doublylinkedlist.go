package doublylinkedlist

type edges struct {
	first  *doublylinkedlist
	last   *doublylinkedlist
	length int
}

type doublylinkedlist struct {
	payload interface{}
	edges   *edges
	next    *doublylinkedlist
	prev    *doublylinkedlist
}

type predicate func(item DoublyLinkedList, index int) bool

// DoublyLinkedList whatever
type DoublyLinkedList interface {
	Value() interface{}
	Next() DoublyLinkedList
	First() DoublyLinkedList
	Last() DoublyLinkedList
	Insert(payload interface{}) DoublyLinkedList
	Delete() DoublyLinkedList
	Update(payload interface{}) DoublyLinkedList
	Length() int

	Filter(p predicate) DoublyLinkedList
}

func isEmptyList(l *doublylinkedlist) bool {
	return l.edges.first.Next() == nil
}

func isFirstInList(l *doublylinkedlist) bool {
	return l.edges.first == l
}

func (l *doublylinkedlist) Filter(p predicate) DoublyLinkedList {
	ll := New()
	i := 0
	for curr := l.edges.first.next; curr != nil && i < l.edges.length; curr = curr.next {
		if p(curr, i) {
			ll.Last().Insert(curr.payload)
		}
		i++
	}
	return ll
}

func (l *doublylinkedlist) Length() int {
	return l.edges.length
}

func (l *doublylinkedlist) Value() interface{} {
	return l.payload
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
		payload: payload,
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

func (l *doublylinkedlist) Update(payload interface{}) DoublyLinkedList {
	if isEmptyList(l) {
		return l.Insert(payload)
	}
	if isFirstInList(l) {
		l.next.payload = payload
		return l.next
	}
	l.payload = payload
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

package bstree

import (
	"github.com/rianby64/data-structures-self-study/cell"
	"github.com/rianby64/data-structures-self-study/list"
)

// comparator type.
type comparator func(a, b interface{}) bool

// BStree stands for Binary Search Tree interface.
type BStree interface {
	Value() interface{}
	Insert(v interface{}) BStree
	Delete() BStree

	Inorder() list.List
	Length() int
	Parent() BStree
	Left() BStree
	Right() BStree

	Find(value interface{}, comparator comparator) BStree
}

type bstree struct {
	payload    cell.Cell
	comparator comparator
	root       *bstree
	left       *bstree
	right      *bstree
	parent     *bstree
	length     int
}

func (t *bstree) Delete() BStree {
	if t.left == nil && t.right == nil && t.parent == nil {
		t.payload = nil

		return nil
	}

	return delete(t, t.parent)
}

func (t *bstree) Length() int {
	return t.length
}

func (t *bstree) Parent() BStree {
	if t.parent == nil {
		return nil
	}

	return t.parent
}

func (t *bstree) Left() BStree {
	if t.left == nil {
		return nil
	}

	return t.left
}

func (t *bstree) Right() BStree {
	if t.right == nil {
		return nil
	}

	return t.right
}

func (t *bstree) Find(value interface{}, matcher comparator) BStree {
	if matcher == nil {
		return nil
	}

	return find(value, t, matcher)
}

func (t *bstree) Value() interface{} {
	if t.payload != nil {
		return t.payload.Value()
	}

	return nil
}

func (t *bstree) Inorder() list.List {
	root := t.root
	list := list.New()

	inorder(root, list)

	return list
}

func (t *bstree) Insert(v interface{}) BStree {
	root := t.root

	return insert(root, v, root.comparator)
}

// New constructor.
func New(comparator func(a, b interface{}) bool) BStree {
	t := &bstree{}
	t.root = t
	t.parent = nil
	t.comparator = func(a, b interface{}) bool {
		if comparator == nil {
			return false
		}

		return comparator(a, b)
	}

	return t
}

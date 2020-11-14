package bstree

import (
	"github.com/rianby64/data-structures-self-study/cell"
	"github.com/rianby64/data-structures-self-study/list"
)

// comparator type
type comparator func(a, b interface{}) bool

// BStree stands for Binary Search Tree interface
type BStree interface {
	cell.Cell
	Insert(v interface{})
	Inorder() list.List
}

type bstree struct {
	payload    cell.Cell
	comparator comparator
	root       *bstree
	left       *bstree
	right      *bstree
}

func (t *bstree) Value() interface{} {
	if t.payload != nil {
		return t.payload.Value()
	}
	return nil
}

func (t *bstree) SetValue(v interface{}) {
	if t.payload != nil {
		t.payload.SetValue(v)
	}
	t.payload = cell.New(v)
}

func inorder(root *bstree, l list.List) {
	if root == nil {
		return
	}

	if root.left != nil {
		inorder(root.left, l)
	}

	l.Last().Insert(root.Value())

	if root.right != nil {
		inorder(root.right, l)
	}
}

func (t *bstree) Inorder() list.List {
	root := t.root
	list := list.New()

	inorder(root, list)

	return list
}

func insert(t *bstree, v interface{}, c comparator) {
	if t.payload == nil {
		t.SetValue(v)
		return
	}

	if c(t.payload.Value(), v) {
		if t.left == nil {
			t.left = &bstree{
				root:       t.root,
				comparator: c,
			}
		}
		insert(t.left, v, c)
	} else {
		if t.right == nil {
			t.right = &bstree{
				root:       t.root,
				comparator: c,
			}
		}
		insert(t.right, v, c)
	}
}

func (t *bstree) Insert(v interface{}) {
	root := t.root

	insert(root, v, root.comparator)
}

// New constructor
func New(c comparator) BStree {
	t := &bstree{}
	t.root = t
	t.comparator = c
	return t
}

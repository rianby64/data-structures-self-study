package bstree

import (
	"github.com/rianby64/data-structures-self-study/cell"
	"github.com/rianby64/data-structures-self-study/list"
)

// BStree stands for Binary Search Tree interface
type BStree interface {
	cell.Cell
	Insert(v interface{})
	Inorder() list.List
}

type bstree struct {
	payload cell.Cell
	root    *bstree
	left    *bstree
	right   *bstree
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

func insert(t *bstree, v interface{}) {
	if t.payload == nil {
		t.SetValue(v)
		return
	}

	// aqui esto es motivo de llevarlo a una funcion comparadora
	vnode := t.payload.Value().(int)
	vinse := v.(int)

	if vnode > vinse {
		if t.left == nil {
			t.left = &bstree{
				root: t.root,
			}
		}
		insert(t.left, v)
	} else {
		if t.right == nil {
			t.right = &bstree{
				root: t.root,
			}
		}
		insert(t.right, v)
	}
}

func (t *bstree) Insert(v interface{}) {
	root := t.root

	insert(root, v)
}

// New constructor
func New() BStree {
	t := &bstree{}
	t.root = t
	return t
}

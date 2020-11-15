package bstree

import (
	"testing"

	"github.com/rianby64/data-structures-self-study/list"
	"github.com/stretchr/testify/assert"
)

func checkExpected(ll list.List, expected []int, t *testing.T) {
	actual := []int{}
	i := 0
	for curr := ll.Next(); curr != nil; curr = curr.Next() {
		if i > len(expected) {
			break
		}
		actual = append(actual, curr.Value().(int))
		i++
	}

	assert.Equal(t, i, len(expected))
	assert.Equal(t, len(expected), ll.Length())
	assert.Equal(t, expected, actual)
}

func Test_tree_inorder_comparator_nil(t *testing.T) {
	btree := New(nil)

	n1 := btree.Insert(33)
	n2 := btree.Insert(44)
	n3 := btree.Insert(55)

	expectedNodes := []BStree{n1, n2, n3}
	l := btree.Inorder()

	i := 0
	for curr := l.Next(); curr != nil; curr = curr.Next() {
		expectedNode := expectedNodes[i]
		i++
		assert.Equal(t, expectedNode.Value(), curr.Value())
	}
}

func Test_tree_inorder_empty(t *testing.T) {
	btree := New(nil)

	l := btree.Inorder()
	assert.Zero(t, l.Length())
	assert.Nil(t, btree.Parent())
	assert.Nil(t, btree.Left())
	assert.Nil(t, btree.Right())
}

func Test_tree_case_inorder(t *testing.T) {
	c := func(a, b interface{}) bool {
		return a.(int) >= b.(int)
	}
	btree := New(c)

	btree.Insert(50)
	btree.Insert(30)
	btree.Insert(20)
	btree.Insert(40)
	btree.Insert(70)
	btree.Insert(60)
	btree.Insert(80)

	expected := []int{20, 30, 40, 50, 60, 70, 80}

	l := btree.Inorder()
	checkExpected(l, expected, t)
}

func Test_tree_case_inorder_chain(t *testing.T) {
	c := func(a, b interface{}) bool {
		return a.(int) >= b.(int)
	}
	btree := New(c)

	item := btree.Insert(50).Insert(30)
	btree.Insert(20).Insert(40).Insert(70).Insert(60).Insert(80)

	expected := []int{20, 30, 40, 50, 60, 70, 80}

	l := btree.Inorder()
	checkExpected(l, expected, t)

	assert.Equal(t, item.Value(), 30)
	assert.Equal(t, item.Parent().Value(), 50)
	assert.Equal(t, item.Left().Value(), 20)
	assert.Equal(t, item.Right().Value(), 40)
}

func Test_tree_search_ok_middle_case_1(t *testing.T) {
	c := func(a, b interface{}) bool {
		na := a.(int)
		nb := b.(int)
		return na >= nb
	}
	btree := New(c)

	expected := btree.Insert(50).Insert(30)
	btree.Insert(20).Insert(40).Insert(70).Insert(60).Insert(80)

	actual := btree.Find(30, func(a, b interface{}) bool {
		na := a.(int)
		nb := b.(int)

		return na == nb
	})

	assert.Equal(t, expected, actual)
}

func Test_tree_search_ok_middle_case_2(t *testing.T) {
	c := func(a, b interface{}) bool {
		na := a.(int)
		nb := b.(int)
		return na >= nb
	}
	btree := New(c)

	expected := btree.Insert(50).Insert(30).Insert(20).Insert(40)
	btree.Insert(70).Insert(60).Insert(80)

	actual := btree.Find(40, func(a, b interface{}) bool {
		na := a.(int)
		nb := b.(int)

		return na == nb
	})

	assert.Equal(t, expected, actual)
}

func Test_tree_search_ok_first(t *testing.T) {
	c := func(a, b interface{}) bool {
		na := a.(int)
		nb := b.(int)
		return na >= nb
	}
	btree := New(c)

	expected := btree.Insert(50)
	btree.Insert(30).Insert(20).Insert(40).Insert(70).Insert(60).Insert(80)

	actual := btree.Find(50, func(a, b interface{}) bool {
		na := a.(int)
		nb := b.(int)

		return na == nb
	})

	assert.Equal(t, expected, actual)
}

func Test_tree_search_ok_last(t *testing.T) {
	c := func(a, b interface{}) bool {
		na := a.(int)
		nb := b.(int)
		return na >= nb
	}
	btree := New(c)

	expected := btree.Insert(50).Insert(30).Insert(20).Insert(40).Insert(70).Insert(60).Insert(80)

	actual := btree.Find(80, func(a, b interface{}) bool {
		na := a.(int)
		nb := b.(int)

		return na == nb
	})

	assert.Equal(t, expected, actual)
}

func Test_tree_search_nil(t *testing.T) {
	c := func(a, b interface{}) bool {
		na := a.(int)
		nb := b.(int)
		return na >= nb
	}
	btree := New(c)

	btree.Insert(50).Insert(30).Insert(20).Insert(40).Insert(70).Insert(60).Insert(80)

	actual := btree.Find(100, func(a, b interface{}) bool {
		na := a.(int)
		nb := b.(int)

		return na == nb
	})

	assert.Equal(t, nil, actual)
}

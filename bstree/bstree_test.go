package bstree

import (
	"testing"

	"github.com/rianby64/data-structures-self-study/list"
	"github.com/stretchr/testify/assert"
)

func castTobtree(b BStree) (*bstree, bool) {
	casted, ok := b.(*bstree)
	return casted, ok
}

func corder(a, b interface{}) bool {
	na := a.(int)
	nb := b.(int)
	return na >= nb
}

func cequal(a, b interface{}) bool {
	na := a.(int)
	nb := b.(int)

	return na == nb
}

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
	btree := New(corder)

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
	btree := New(corder)

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

func Test_tree_search_comparator_nil(t *testing.T) {
	btree := New(corder)

	btree.Insert(50).Insert(30).Insert(20).Insert(40).Insert(70).Insert(60).Insert(80)

	actual := btree.Find(30, nil)

	assert.Nil(t, actual)
}

func Test_tree_search_ok_middle_case_1(t *testing.T) {
	btree := New(corder)

	expected := btree.Insert(50).Insert(30)
	btree.Insert(20).Insert(40).Insert(70).Insert(60).Insert(80)

	actual := btree.Find(30, cequal)

	assert.Equal(t, expected, actual)
}

func Test_tree_search_ok_middle_case_2(t *testing.T) {
	btree := New(corder)

	expected := btree.Insert(50).Insert(30).Insert(20).Insert(40)
	btree.Insert(70).Insert(60).Insert(80)

	actual := btree.Find(40, cequal)

	assert.Equal(t, expected, actual)
}

func Test_tree_search_ok_first(t *testing.T) {
	btree := New(corder)

	expected := btree.Insert(50)
	btree.Insert(30).Insert(20).Insert(40).Insert(70).Insert(60).Insert(80)

	actual := btree.Find(50, cequal)

	assert.Equal(t, expected, actual)
}

func Test_tree_search_ok_last(t *testing.T) {
	btree := New(corder)

	expected := btree.Insert(50).Insert(30).Insert(20).Insert(40).Insert(70).Insert(60).Insert(80)

	actual := btree.Find(80, cequal)

	assert.Equal(t, expected, actual)
}

func Test_tree_length_empty(t *testing.T) {
	btree := New(corder)

	assert.Equal(t, 0, btree.Length())
}

func Test_tree_length_one_leaf(t *testing.T) {
	btree := New(corder)

	btree.Insert(50)

	assert.Equal(t, 1, btree.Length())
}

func Test_tree_length_many_items(t *testing.T) {
	btree := New(corder)

	actualNodes := []BStree{
		btree.Insert(50),
		btree.Insert(30),
		btree.Insert(20),
		btree.Insert(40),
		btree.Insert(70),
		btree.Insert(60),
		btree.Insert(80),
	}

	expectedLenghts := []int{7, 3, 1, 1, 3, 1, 1}
	for i, expectedLength := range expectedLenghts {
		assert.Equal(t, expectedLength, actualNodes[i].Length(), i)
	}

	assert.Equal(t, 7, btree.Length())
}

func Test_tree_search_nil(t *testing.T) {
	btree := New(corder)

	btree.Insert(50).Insert(30).Insert(20).Insert(40).Insert(70).Insert(60).Insert(80)

	actual := btree.Find(100, cequal)

	assert.Equal(t, nil, actual)
}

func Test_tree_delete_emtpy(t *testing.T) {
	btree := New(corder)

	deleted := btree.Delete()
	assert.Nil(t, deleted)
}

func Test_tree_delete_case_1(t *testing.T) {
	btree := New(corder)

	expected := []int{20, 30, 40, 50, 60, 70}
	last := btree.Insert(50).Insert(30).Insert(20).Insert(40).Insert(70).Insert(60).Insert(80)

	last.Delete()

	l := btree.Inorder()
	checkExpected(l, expected, t)
}

func Test_tree_delete_case_2(t *testing.T) {
	btree := New(corder)
	toDelete := btree.Insert(50).Insert(30).Insert(20).Insert(40).Insert(70)
	btree.Insert(80).Insert(60).Insert(65).Insert(55)

	expected := []int{20, 30, 40, 50, 55, 60, 65, 80}

	btreeExpected := New(corder)
	btreeExpected.Insert(50).Insert(30).Insert(20).Insert(40).Insert(65).Insert(80).Insert(60).Insert(55)

	toDelete.Delete()

	l := btree.Inorder()
	checkExpected(l, expected, t)

	lexpected := btreeExpected.Inorder()
	checkExpected(lexpected, expected, t)
}

func Test_tree_delete_case_3(t *testing.T) {
	btree := New(corder)
	toDelete := btree.Insert(50)
	btree.Insert(30).Insert(20).Insert(40).Insert(70).Insert(80).Insert(60).Insert(65).Insert(55)

	expected := []int{20, 30, 40, 55, 60, 65, 70, 80}

	btreeExpected := New(corder)
	btreeExpected.Insert(20).Insert(30).Insert(40).Insert(70).Insert(80).Insert(60).Insert(65).Insert(55)

	toDelete.Delete()

	l := btree.Inorder()
	checkExpected(l, expected, t)

	lexpected := btreeExpected.Inorder()
	checkExpected(lexpected, expected, t)
}

func Test_tree_delete_case_4(t *testing.T) {
	btree := New(corder)
	deleted := btree.Insert(50).Insert(30).Insert(20).Insert(40).Insert(70)
	btree.Insert(80).Insert(60).Insert(65).Insert(55)

	{
		expected := []int{20, 30, 40, 50, 55, 60, 65, 80}

		btreeExpected := New(corder)
		btreeExpected.Insert(50).Insert(20).Insert(30).Insert(40).Insert(65).Insert(80).Insert(60).Insert(55)

		deleted = deleted.Delete()
		assert.Equal(t, 65, deleted.Value())

		l := btree.Inorder()
		checkExpected(l, expected, t)

		lexpected := btreeExpected.Inorder()
		checkExpected(lexpected, expected, t)
	}

	{
		expected := []int{20, 30, 40, 50, 55, 60, 80}

		btreeExpected := New(corder)
		btreeExpected.Insert(50).Insert(20).Insert(30).Insert(40).Insert(60).Insert(80).Insert(55)

		deleted = deleted.Delete()
		assert.Equal(t, 60, deleted.Value())

		l := btree.Inorder()
		checkExpected(l, expected, t)

		lexpected := btreeExpected.Inorder()
		checkExpected(lexpected, expected, t)
	}

	{
		expected := []int{20, 30, 40, 50, 55, 80}

		btreeExpected := New(corder)
		btreeExpected.Insert(50).Insert(20).Insert(30).Insert(40).Insert(55).Insert(80)

		deleted = deleted.Delete()
		assert.Equal(t, 55, deleted.Value())

		l := btree.Inorder()
		checkExpected(l, expected, t)

		lexpected := btreeExpected.Inorder()
		checkExpected(lexpected, expected, t)
	}

	{
		expected := []int{20, 30, 40, 50, 80}

		btreeExpected := New(corder)
		btreeExpected.Insert(50).Insert(20).Insert(30).Insert(40).Insert(80)

		deleted = deleted.Delete()
		assert.Equal(t, 80, deleted.Value())

		l := btree.Inorder()
		checkExpected(l, expected, t)

		lexpected := btreeExpected.Inorder()
		checkExpected(lexpected, expected, t)
	}

	{
		expected := []int{20, 30, 40, 50}

		btreeExpected := New(corder)
		btreeExpected.Insert(50).Insert(20).Insert(30).Insert(40)

		deleted = deleted.Delete()
		assert.Equal(t, 50, deleted.Value())

		l := btree.Inorder()
		checkExpected(l, expected, t)

		lexpected := btreeExpected.Inorder()
		checkExpected(lexpected, expected, t)
	}

	{
		expected := []int{20, 30, 40}

		btreeExpected := New(corder)
		btreeExpected.Insert(30).Insert(20).Insert(40)

		deleted = deleted.Delete()
		assert.Equal(t, 40, deleted.Value())

		l := btree.Inorder()
		checkExpected(l, expected, t)

		lexpected := btreeExpected.Inorder()
		checkExpected(lexpected, expected, t)
	}

	{
		expected := []int{20, 30}

		btreeExpected := New(corder)
		btreeExpected.Insert(30).Insert(20)

		deleted = deleted.Delete()
		assert.Equal(t, 30, deleted.Value())

		l := btree.Inorder()
		checkExpected(l, expected, t)

		lexpected := btreeExpected.Inorder()
		checkExpected(lexpected, expected, t)
	}

	{
		expected := []int{20}

		btreeExpected := New(corder)
		btreeExpected.Insert(20)

		deleted = deleted.Delete()
		assert.Equal(t, 20, deleted.Value())

		l := btree.Inorder()
		checkExpected(l, expected, t)

		lexpected := btreeExpected.Inorder()
		checkExpected(lexpected, expected, t)
	}

	{
		deleted = deleted.Delete()
		assert.Nil(t, deleted)

		assert.Nil(t, btree.Value())
	}
}

func Test_tree_findmax_case_1(t *testing.T) {
	btree := New(corder)

	btree.Insert(50).Insert(30).Insert(40).Insert(70).Insert(80).Insert(20)

	expected := []int{20, 30, 40, 50, 70, 80}

	l := btree.Inorder()
	checkExpected(l, expected, t)

	cbtree, _ := castTobtree(btree)
	max := findmax(cbtree)
	assert.Equal(t, 80, max.Value())
}

func Test_tree_findmax_case_2(t *testing.T) {
	btree := New(corder)

	btree.Insert(50).Insert(30).Insert(40).Insert(70).Insert(20).Insert(60).Insert(55).Insert(65)

	expected := []int{20, 30, 40, 50, 55, 60, 65, 70}

	l := btree.Inorder()
	checkExpected(l, expected, t)

	cbtree, _ := castTobtree(btree)
	max := findmax(cbtree)
	assert.Equal(t, 70, max.Value())
}

func Test_tree_findmax_emtpy_tree(t *testing.T) {
	btree := New(corder)

	cbtree, _ := castTobtree(btree)
	max := findmax(cbtree)
	assert.Nil(t, max.Value())
}

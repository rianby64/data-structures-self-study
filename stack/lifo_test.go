package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_empty(t *testing.T) {
	stack := New()

	pnil := stack.Pop()
	assert.Nil(t, pnil)

	pnil2 := stack.Pop()
	assert.Nil(t, pnil2)
}

func Test_pushOne_popOne(t *testing.T) {
	stack := New()
	stack.Push(33)

	p1 := stack.Pop()
	assert.NotNil(t, p1)
	assert.Equal(t, p1.(int), 33)

	pnil := stack.Pop()
	assert.Nil(t, pnil)

	pnil2 := stack.Pop()
	assert.Nil(t, pnil2)
}

func Test_pushTwo_popTwo(t *testing.T) {
	stack := New()
	stack.Push(33)
	stack.Push(44)

	p2 := stack.Pop()
	assert.NotNil(t, p2)
	assert.Equal(t, p2.(int), 44)

	p1 := stack.Pop()
	assert.NotNil(t, p1)
	assert.Equal(t, p1.(int), 33)

	pnil := stack.Pop()
	assert.Nil(t, pnil)
}

func Test_pushThree_popThree(t *testing.T) {
	stack := New()

	stack.Push(33)
	stack.Push(44)
	stack.Push(55)

	p3 := stack.Pop()
	assert.NotNil(t, p3)
	assert.Equal(t, p3.(int), 55)

	p2 := stack.Pop()
	assert.NotNil(t, p2)
	assert.Equal(t, p2.(int), 44)

	p1 := stack.Pop()
	assert.NotNil(t, p1)
	assert.Equal(t, p1.(int), 33)

	pnil := stack.Pop()
	assert.Nil(t, pnil)
}

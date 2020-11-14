package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_empty_dequeue(t *testing.T) {
	q := New()

	v := q.Dequeue()
	assert.Nil(t, v)
}

func Test_queue_enqueueThreeItems_dequeueFourItems(t *testing.T) {
	q := New()

	q.Enqueue(22)
	q.Enqueue(33)
	q.Enqueue(44)

	v1 := q.Dequeue()
	assert.Equal(t, v1, 22)

	v2 := q.Dequeue()
	assert.Equal(t, v2, 33)

	v3 := q.Dequeue()
	assert.Equal(t, v3, 44)

	v4 := q.Dequeue()
	assert.Nil(t, v4)
}

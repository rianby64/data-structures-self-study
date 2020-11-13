package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cell(t *testing.T) {
	c := New(33)
	v := c.Value()
	assert.NotNil(t, v)
	assert.Equal(t, v.(int), 33)

	c.SetValue(55)
	u := c.Value()
	assert.NotNil(t, u)
	assert.Equal(t, u.(int), 55)
}

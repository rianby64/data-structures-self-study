package dict

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Dict_Get(t *testing.T) {
	d := New()

	assert.Nil(t, d.Get(""))
	assert.Nil(t, d.Get("a"))
	assert.Nil(t, d.Get("abc"))
}

func Test_Dict_Set(t *testing.T) {
	d := New()

	assert.True(t, d.Set("abc", 33))
	assert.Equal(t, 33, d.Get("abc"))

	assert.True(t, d.Set("ab", 44))
	assert.Equal(t, 44, d.Get("ab"))

	assert.True(t, d.Set("ba", 55))
	assert.Equal(t, 55, d.Get("ba"))

	assert.True(t, d.Set("b", 66))
	assert.Equal(t, 66, d.Get("b"))

	assert.True(t, d.Set("ac", 87))
	assert.Equal(t, 87, d.Get("ac"))

	assert.True(t, d.Set("ad", 88))
	assert.Equal(t, 88, d.Get("ad"))

	assert.True(t, d.Set("ac", 187))
	assert.Equal(t, 187, d.Get("ac"))

	assert.Nil(t, d.Get("abcd"))
	assert.Nil(t, d.Get("azcd"))
}

func Test_Dict_Delete(t *testing.T) {
	d := New()

	assert.False(t, d.Set("", nil))
	assert.True(t, d.Set("abc", 55))
	assert.True(t, d.Del("abc"))
	assert.Nil(t, d.Get("abc"))
	assert.True(t, d.Set("abc", 55))
	assert.Equal(t, 55, d.Get("abc"))
}

func Test_Dict_heavy(t *testing.T) {
	d := New()

	for i := 0; i < 1000; i++ {
		r, err := uuid.NewRandom()
		if err != nil {
			t.Error(err)
			return
		}

		key := r.String()

		assert.Nil(t, d.Get(key), key)

		d.Set(key, true)

		assert.Equal(t, true, d.Get(key))
	}
}

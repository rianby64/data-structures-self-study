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

	assert.False(t, d.Set("g", nil))
	assert.Nil(t, d.Get("g"))

	assert.True(t, d.Set("abc", 33))
	assert.True(t, d.Set("ab", 44))
	assert.True(t, d.Set("ba", 55))
	assert.True(t, d.Set("b", 66))
	assert.True(t, d.Set("ac", 87))
	assert.Equal(t, 87, d.Get("ac"))

	assert.True(t, d.Set("ad", 88))
	assert.True(t, d.Set("ac", 187))

	assert.Equal(t, 33, d.Get("abc"))
	assert.Equal(t, 44, d.Get("ab"))
	assert.Equal(t, 55, d.Get("ba"))
	assert.Equal(t, 66, d.Get("b"))
	assert.Equal(t, 88, d.Get("ad"))
	assert.Equal(t, 187, d.Get("ac"))

	assert.Nil(t, d.Get("abcd"))
	assert.Nil(t, d.Get("azcd"))
}

func Test_Dict_Delete(t *testing.T) {
	d := New()

	assert.False(t, d.Set("", nil))
	assert.False(t, d.Del(""))
	assert.False(t, d.Del("z"))
	assert.True(t, d.Set("abc", 55))
	assert.True(t, d.Del("abc"))
	assert.Nil(t, d.Get("abc"))
	assert.True(t, d.Set("abc", 55))
	assert.Equal(t, 55, d.Get("abc"))
}

func Test_Dict_ogo(t *testing.T) {
	d := New()

	d.Set("edac6cfd-50b3-42e2-978f-92f4810d52f2", true)
	d.Set("5f69b8ba-3ddc-4029-90f0-f6e7f482878b", true)
}

func Test_Dict_heavy(t *testing.T) {
	d := New()

	for i := 0; i < 10000; i++ {
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

func Test_Dict_Delete_cleanup_case1(t *testing.T) {
	d := &dict{
		payload: newLevel(),
	}

	assert.True(t, d.Set("abc", 55))
	assert.True(t, d.Del("abc"))
	assert.Nil(t, d.Get("abc"))

	assert.Zero(t, len(d.payload.payload))
}

func Test_Dict_Delete_cleanup_case2(t *testing.T) {
	d := &dict{
		payload: newLevel(),
	}

	assert.True(t, d.Set("ab", 44))
	assert.True(t, d.Set("abc", 55))

	assert.True(t, d.Del("abc"))
	assert.Nil(t, d.Get("abc"))
	assert.Equal(t, 44, d.Get("ab"))
	assert.Equal(t, 1, len(d.payload.payload))

	assert.True(t, d.Del("ab"))
	assert.Nil(t, d.Get("ab"))
	assert.Equal(t, 0, len(d.payload.payload))
}

func Test_Dict_Delete_cleanup_case3(t *testing.T) {
	d := &dict{
		payload: newLevel(),
	}

	assert.True(t, d.Set("ab", 44))
	assert.True(t, d.Set("abc", 55))

	assert.False(t, d.Del("a"))

	assert.True(t, d.Del("abc"))
	assert.Nil(t, d.Get("abc"))
	assert.Equal(t, 44, d.Get("ab"))
	assert.Equal(t, 1, len(d.payload.payload))

	assert.True(t, d.Del("ab"))
	assert.Nil(t, d.Get("ab"))
	assert.Equal(t, 0, len(d.payload.payload))
}

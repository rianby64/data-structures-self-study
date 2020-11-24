package dict

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	key      byte
	expected int
}

func Test_Level_insert_incr(t *testing.T) {

	d := newLevel()

	d.insert(97, nil)
	d.insert(98, nil)
	d.insert(99, nil)

	assert.Equal(t, "abc", d.getkeys())
}

func Test_Level_insert_decr(t *testing.T) {

	d := newLevel()

	d.insert(99, nil)
	d.insert(98, nil)
	d.insert(97, nil)

	assert.Equal(t, "abc", d.getkeys())
}

func Test_Level_insert_incr_repeated(t *testing.T) {

	d := newLevel()

	d.insert(97, nil)
	d.insert(98, nil)
	d.insert(99, nil)

	d.insert(97, nil)

	assert.Equal(t, "abc", d.getkeys())
}

func Test_Level_insert_decr_repeated(t *testing.T) {

	d := newLevel()

	d.insert(99, nil)
	d.insert(98, nil)
	d.insert(97, nil)

	d.insert(99, nil)

	assert.Equal(t, "abc", d.getkeys())
}

func Test_Level_insert_in_the_middle(t *testing.T) {

	d := newLevel()

	d.insert(97, nil)
	d.insert(122, nil)

	d.insert(104, nil)

	assert.Equal(t, "ahz", d.getkeys())
}

func Test_Level_insert_in_the_middle_repeated(t *testing.T) {

	d := newLevel()

	d.insert(97, nil)
	d.insert(122, nil)

	d.insert(104, nil)
	d.insert(103, nil)

	d.insert(104, nil)

	assert.Equal(t, "aghz", d.getkeys())
}

func Test_Level_getIndex_case1(t *testing.T) {
	d := newLevel()
	testValues := []testcase{
		{byte(97), 0},
		{byte(122), 5},
		{byte(104), 2},
		{byte(103), 1},
		{byte(115), 4},
		{byte(114), 3},
	}

	for _, testValue := range testValues {
		d.insert(testValue.key, nil)
	}

	assert.Equal(t, "aghrsz", d.getkeys())

	for _, testValue := range testValues {
		v, ok := d.getIndex(testValue.key)
		assert.True(t, ok)
		assert.Equal(t, testValue.expected, v)

	}
}

func Test_Level_getIndex_case2(t *testing.T) {
	d := newLevel()
	testValues := []testcase{}

	for i := 97; i <= 122; i++ {
		tc := testcase{byte(i), i - 97}
		testValues = append(testValues, tc)
	}

	for _, testValue := range testValues {
		d.insert(testValue.key, nil)
	}

	assert.Equal(t, "abcdefghijklmnopqrstuvwxyz", d.getkeys())

	for _, testValue := range testValues {
		v, ok := d.getIndex(testValue.key)
		assert.True(t, ok)
		assert.Equal(t, testValue.expected, v)

	}
}

func Test_Level_getIndex_case3(t *testing.T) {

	d := newLevel()

	d.insert(97, nil)
	d.insert(122, nil)

	assert.Equal(t, "az", d.getkeys())

	i, ok := d.getIndex(98)
	assert.False(t, ok)
	assert.Equal(t, -1, i)
}

func Test_Level_getIndex_case4(t *testing.T) {

	d := newLevel()

	d.insert(97, nil)
	d.insert(122, nil)

	assert.Equal(t, "az", d.getkeys())

	i, ok := d.getIndex(121)
	assert.False(t, ok)
	assert.Equal(t, -1, i)
}

func Test_Level_delete_min(t *testing.T) {

	d := newLevel()

	d.insert(97, nil)
	d.insert(122, nil)
	d.insert(104, nil)

	assert.Equal(t, "ahz", d.getkeys())

	d.delete(97)

	assert.Equal(t, "hz", d.getkeys())
}

func Test_Level_delete_max(t *testing.T) {

	d := newLevel()

	d.insert(97, nil)
	d.insert(122, nil)
	d.insert(104, nil)

	assert.Equal(t, "ahz", d.getkeys())

	d.delete(122)

	assert.Equal(t, "ah", d.getkeys())
}

func Test_Level_delete_middle(t *testing.T) {

	d := newLevel()

	d.insert(97, nil)
	d.insert(122, nil)
	d.insert(104, nil)

	assert.Equal(t, "ahz", d.getkeys())

	d.delete(104)

	assert.Equal(t, "az", d.getkeys())
}

/*
func Test_Dict(t *testing.T) {

	d := New()

	r := d.Get("")
	assert.Nil(t, r)

}

func Test_Dict(t *testing.T) {

	d := New()

	d.Set("key", 33)
	assert.Equal(t, 33, d.Get("key"))

}
*/

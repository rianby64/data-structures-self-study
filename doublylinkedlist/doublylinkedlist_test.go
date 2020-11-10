package doublylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkExpected(ll DoublyLinkedList, expected []int, t *testing.T) {
	actual := []int{}
	i := 0
	for curr := ll.Next(); curr != nil; curr = curr.Next() {
		if i > len(expected) {
			break
		}
		actual = append(actual, curr.Value().(int))
		assert.Equal(t, expected[0], curr.First().Value().(int))
		assert.Equal(t, expected[len(expected)-1], curr.Last().Value().(int))
		i++
	}

	assert.Equal(t, i, len(expected))
	assert.Equal(t, len(expected), ll.Length())
	assert.Equal(t, expected, actual)
}

func Test_emptyList(t *testing.T) {
	ll := New()
	i := 0

	first := ll.First()
	last := ll.Last()

	assert.NotNil(t, first)
	assert.NotNil(t, last)

	assert.Nil(t, first.Value())
	assert.Nil(t, last.Value())

	for curr := ll.Next(); curr != nil && i < 2; curr = curr.Next() {
		i++
	}

	assert.Equal(t, 0, i)
	assert.Nil(t, ll.Value())
}

func Test_addOnePayload(t *testing.T) {
	ll := New()
	expected := 33

	ll.Insert(expected)
	actual := ll.First().Value().(int)

	i := 0
	for curr := ll.Next(); curr != nil && i < 2; curr = curr.Next() {
		i++
	}

	first := ll.First()
	last := ll.Last()

	assert.NotNil(t, first)
	assert.NotNil(t, last)

	assert.Equal(t, expected, first.Value())
	assert.Equal(t, expected, last.Value())

	assert.Equal(t, actual, expected)
	assert.Equal(t, 1, i)
}

func Test_addTwoPayloads(t *testing.T) {
	ll := New()

	expected := []int{33, 44}
	ll.Insert(expected[0]).Insert(expected[1])
	checkExpected(ll, expected, t)
}

func Test_addThreePayloads(t *testing.T) {
	ll := New()

	expected := []int{33, 44, 55}
	ll.Insert(expected[0]).Insert(expected[1]).Insert(expected[2])
	checkExpected(ll, expected, t)
}

func Test_addTwoPayloadsFromRoot(t *testing.T) {
	ll := New()

	expected := []int{33, 44}
	ll.Insert(expected[1])
	ll.Insert(expected[0])
	checkExpected(ll, expected, t)
}

func Test_addThreePayloadsFromRoot(t *testing.T) {
	ll := New()

	expected := []int{33, 44, 55}
	ll.Insert(expected[2])
	ll.Insert(expected[1])
	ll.Insert(expected[0])
	checkExpected(ll, expected, t)
}

func Test_addThreePayloadsThenFromFirstAddOne(t *testing.T) {
	ll := New()

	expected := []int{33, 44, 55, 66}
	ll.Insert(expected[0]).Insert(expected[2]).Insert(expected[3])
	ll.First().Insert(expected[1])
	checkExpected(ll, expected, t)
}

func Test_addThreePayloadsThenFromSecondAddOne(t *testing.T) {
	ll := New()

	expected := []int{33, 44, 55, 66}
	ll.Insert(expected[0]).Insert(expected[1]).Insert(expected[3])
	ll.First().Next().Insert(expected[2])
	checkExpected(ll, expected, t)
}

func Test_addThreePayloadsThenFromLastAddOne(t *testing.T) {
	ll := New()

	expected := []int{33, 44, 55, 66}
	ll.Insert(expected[0]).Insert(expected[1]).Insert(expected[2])
	ll.Last().Insert(expected[3])
	checkExpected(ll, expected, t)
}

func Test_fromEmptyListDelete(t *testing.T) {
	ll := New()

	ll.Delete()

	assert.Equal(t, nil, ll.First().Value())
	assert.Equal(t, nil, ll.Last().Value())
}

func Test_addOnePayloadThenDeleteFromRoot(t *testing.T) {
	ll := New()

	ll.Insert(33)
	ll.Delete()

	assert.Equal(t, nil, ll.First().Value())
	assert.Equal(t, nil, ll.Last().Value())
}

func Test_addOnePayloadThenDeleteFirst(t *testing.T) {
	ll := New()

	ll.Insert(33)
	ll.First().Delete()

	assert.Equal(t, nil, ll.First().Value())
	assert.Equal(t, nil, ll.Last().Value())
}

func Test_addOnePayloadThenDeleteLast(t *testing.T) {
	ll := New()

	ll.Insert(33)
	ll.First().Delete()

	assert.Equal(t, nil, ll.First().Value())
	assert.Equal(t, nil, ll.Last().Value())
}

func Test_addTwoPayloadsThenDeleteFromRoot(t *testing.T) {
	ll := New()

	expected := []int{44}
	ll.Insert(33).Insert(expected[0])
	ll.Delete()

	checkExpected(ll, expected, t)
}

func Test_addTwoPayloadsThenDeleteFirst(t *testing.T) {
	ll := New()

	expected := []int{44}
	ll.Insert(33).Insert(expected[0])
	ll.First().Delete()

	checkExpected(ll, expected, t)
}

func Test_addTwoPayloadsThenDeleteLast(t *testing.T) {
	ll := New()

	expected := []int{44}
	ll.Insert(33).Insert(expected[0])
	ll.First().Delete()

	checkExpected(ll, expected, t)
}

func Test_addThreePayloadsThenDeleteFirst(t *testing.T) {
	ll := New()

	expected := []int{33, 44, 55, 66}
	ll.Insert(100).Insert(expected[0]).Insert(expected[1]).Insert(expected[2]).Insert(expected[3])
	ll.First().Delete()

	checkExpected(ll, expected, t)
}

func Test_addThreePayloadsThenDeleteFromHead(t *testing.T) {
	ll := New()

	expected := []int{33, 44, 55, 66}
	ll.Insert(100).Insert(expected[0]).Insert(expected[1]).Insert(expected[2]).Insert(expected[3])
	ll.Delete()

	checkExpected(ll, expected, t)
}

func Test_addThreePayloadsThenDeleteFromLast(t *testing.T) {
	ll := New()

	expected := []int{33, 44, 55, 66}
	ll.Insert(expected[0]).Insert(expected[1]).Insert(expected[2]).Insert(expected[3]).Insert(100)
	ll.Last().Delete()

	checkExpected(ll, expected, t)
}

func Test_updateEmptyListMeansInsert(t *testing.T) {
	ll := New()

	expected := []int{33}
	ll.Update(expected[0])

	checkExpected(ll, expected, t)
}

func Test_updateEmptyListFromFirstMeansInsert(t *testing.T) {
	ll := New()

	expected := []int{33}
	ll.First().Update(expected[0])

	checkExpected(ll, expected, t)
}

func Test_updateEmptyListFromLastMeansInsert(t *testing.T) {
	ll := New()

	expected := []int{33}
	ll.Last().Update(expected[0])

	checkExpected(ll, expected, t)
}

func Test_addOnePayloadThenUpdateFromRoot(t *testing.T) {
	ll := New()

	expected := []int{33}
	ll.Insert(100)
	ll.Update(expected[0])

	checkExpected(ll, expected, t)
}

func Test_addOnePayloadThenUpdateFromFirst(t *testing.T) {
	ll := New()

	expected := []int{33}
	ll.Insert(100)
	ll.First().Update(expected[0])

	checkExpected(ll, expected, t)
}

func Test_addOnePayloadThenUpdateFromLast(t *testing.T) {
	ll := New()

	expected := []int{33}
	ll.Insert(100)
	ll.Last().Update(expected[0])

	checkExpected(ll, expected, t)
}

func Test_addFivePayloadsThenFilterTrue(t *testing.T) {
	ll := New()

	expected := []int{22, 33, 44, 55, 66, 77, 88, 99}
	curr := ll
	for v := range expected {
		curr = curr.Insert(expected[v])
	}

	filtred := ll.Filter(func(d DoublyLinkedList, i int) bool {
		return true
	})
	checkExpected(ll, expected, t)
	checkExpected(filtred, expected, t)
}

func Test_addFivePayloadsThenFilter445566(t *testing.T) {
	ll := New()

	initial := []int{22, 33, 44, 55, 66, 77, 88, 99}
	expected := []int{44, 55, 66}
	curr := ll
	for v := range initial {
		curr = curr.Insert(initial[v])
	}

	filtred := ll.Filter(func(d DoublyLinkedList, i int) bool {
		n := d.Value().(int)
		return 44 <= n && n <= 66
	})
	checkExpected(ll, initial, t)
	checkExpected(filtred, expected, t)
}

func Test_addFivePayloadsThenFilter445566ByI(t *testing.T) {
	ll := New()

	initial := []int{22, 33, 44, 55, 66, 77, 88, 99}
	expected := []int{44, 55, 66}
	curr := ll
	for v := range initial {
		curr = curr.Insert(initial[v])
	}

	filtred := ll.Filter(func(d DoublyLinkedList, i int) bool {
		return 2 <= i && i <= 4
	})
	checkExpected(ll, initial, t)
	checkExpected(filtred, expected, t)
}

func Test_addFivePayloadsThenFilterToEmpty(t *testing.T) {
	ll := New()

	initial := []int{22, 33, 44, 55, 66, 77, 88, 99}
	expected := []int{}
	curr := ll
	for v := range initial {
		curr = curr.Insert(initial[v])
	}

	filtred := ll.Filter(func(d DoublyLinkedList, i int) bool {
		return false
	})
	checkExpected(ll, initial, t)
	checkExpected(filtred, expected, t)
}

func Test_addFivePayloadsThenFound44(t *testing.T) {
	ll := New()

	initial := []int{22, 33, 44, 55, 66, 77, 88, 99}
	expected := initial[2]
	curr := ll
	for v := range initial {
		curr = curr.Insert(initial[v])
	}

	found := ll.Find(func(d DoublyLinkedList, i int) bool {
		n := d.Value().(int)
		return n == expected
	})
	checkExpected(ll, initial, t)
	assert.Equal(t, expected, found.Value().(int))

	expectedFound := ll.Next().Next().Next()
	assert.Equal(t, &expectedFound, &found)
}

func Test_addFivePayloadsThenFoundNil(t *testing.T) {
	ll := New()

	initial := []int{22, 33, 44, 55, 66, 77, 88, 99}
	curr := ll
	for v := range initial {
		curr = curr.Insert(initial[v])
	}

	found := ll.Find(func(d DoublyLinkedList, i int) bool {
		return false
	})
	checkExpected(ll, initial, t)
	assert.Equal(t, nil, found)
}

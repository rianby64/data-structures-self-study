package dict

import "github.com/rianby64/data-structures-self-study/stack"

type path struct {
	level *level
	key   byte
	index int
	value interface{}
}

type dict struct {
	payload *level
}

// Dict whatever.
type Dict interface {
	Set(string, interface{}) bool
	Del(string) bool
	Get(string) interface{}
	Keys() []string
	Values() []interface{}
}

func (d *dict) Keys() []string {
	return nil
}

func (d *dict) Values() []interface{} {
	return nil
}

func (d *dict) Set(key string, v interface{}) bool {
	if key == "" {
		return false
	}

	if v == nil {
		return false
	}

	bkey := []byte(key)
	l := len(bkey)
	currlevel := d.payload

	for i, b := range bkey {
		value := v
		if i+1 < l {
			value = nil
		}

		c := currlevel.insert(b, value)

		if i+1 < l && c.child == nil {
			c.child = newLevel()
		}

		currlevel = c.child
	}

	return true
}

func cleanup(parent stack.Stack) {
	currparent := parent.Pop()
	for currparent != nil {
		p := currparent.(path)
		cell := p.level.payload[p.index]

		if cell != nil && cell.child != nil && len(cell.child.payload) == 0 {
			p.level.payload[p.index].child = nil
		}

		if p.value == nil {
			p.level.delete(p.key)
		} else {
			break
		}

		currparent = parent.Pop()
	}
}

func (d *dict) Del(key string) bool {
	if key == "" {
		return false
	}

	parent := stack.New()
	bkey := []byte(key)
	l := len(bkey)
	currlevel := d.payload

	defer cleanup(parent)

	for i, b := range bkey {
		k, ok := currlevel.getIndex(b)
		if !ok {
			break
		}

		w := currlevel.payload[k]
		if i+1 == l {
			if w.child == nil {
				return currlevel.delete(b)
			}

			break
		}

		parent.Push(path{currlevel, b, k, w.value})
		currlevel = w.child
	}

	return false
}

func (d *dict) Get(key string) interface{} {
	if key == "" {
		return nil
	}

	bkey := []byte(key)
	l := len(bkey)
	currlevel := d.payload

	for i, b := range bkey {
		k, ok := currlevel.getIndex(b)
		if !ok {
			break
		}

		if i+1 == l {
			return currlevel.payload[k].value
		}

		currlevel = currlevel.payload[k].child
		if currlevel == nil {
			break
		}
	}

	return nil
}

// New constructor.
func New() Dict {
	d := &dict{
		payload: newLevel(),
	}

	return d
}

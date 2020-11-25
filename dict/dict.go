package dict

type dict struct {
	payload *level
}

// Dict whatever.
type Dict interface {
	Set(string, interface{}) bool
	Del(string) bool
	Get(string) interface{}
}

func (d *dict) Set(key string, v interface{}) bool {
	if key == "" {
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

func (d *dict) Del(key string) bool {
	if key == "" {
		return false
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
			return currlevel.delete(b)
		}

		currlevel = currlevel.payload[k].child
		if currlevel == nil {
			break
		}
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

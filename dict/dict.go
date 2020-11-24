package dict

type dict struct {
	payload interface{}
	key     string
	root    *dict
	left    *dict
	right   *dict
	parent  *dict
}

type cell struct {
	key   byte
	value interface{}
}

type level struct {
	changed bool
	min     byte
	max     byte
	len     int
	payload []cell
}

func (l *level) getIndex(b byte) (int, bool) {
	if b < l.min || b > l.max {
		return -1, false
	}

	if b == l.min {
		return 0, true
	}

	if b == l.max {
		return l.len - 1, true
	}

	d := float64(l.max - l.min)

	p := float64(int(b) - int(l.min))
	q := float64(int(l.max) - int(b))

	ip := p / d
	iq := float64(float64(1) - (q / d))

	i := float64(l.len) * (ip + iq) / float64(2)
	index := int(i)
	if index >= l.len {
		index = l.len - 1
	}

	actual := l.payload[index]
	if actual.key == b {
		return index, true
	}

	if actual.key < b {
		for actual.key != b {
			index++
			if index >= l.len {
				return -1, false
			}
			actual = l.payload[index]
		}
		return index, true
	}

	for actual.key != b {
		index--
		if index < 0 {
			return -1, false
		}
		actual = l.payload[index]
	}
	return index, true
}

func (l *level) updateEdges() {
	if !l.changed {
		return
	}
	l.len = len(l.payload)
	l.min = l.payload[0].key
	l.max = l.payload[l.len-1].key
	l.changed = false
}

func (l *level) insert(key byte, value interface{}) {
	defer l.updateEdges()

	if key > l.max {
		l.changed = true
		l.payload = append(l.payload, cell{key: key, value: value})
		return
	}

	if key < l.min {
		l.changed = true
		l.payload = append([]cell{{key: key, value: value}}, l.payload...)
		return
	}

	if key == l.max || key == l.min {
		return
	}

	newpayload := []cell{}
	added := false

	for _, v := range l.payload {
		if v.key == key {
			return
		}

		if v.key > key && !added {
			newpayload = append(newpayload, cell{key: key, value: value})
			added = true
		}
		newpayload = append(newpayload, v)
	}

	l.changed = true
	l.payload = newpayload
}

func (l *level) delete(key byte) {
	if key < l.min || key > l.max {
		return
	}

	defer l.updateEdges()

	newpayload := []cell{}
	changed := false

	for i := 0; i < l.len; i++ {
		v := l.payload[i]
		if v.key == key {
			changed = true
			continue
		}

		newpayload = append(newpayload, v)
	}

	l.changed = changed
	l.payload = newpayload
}

func (l *level) getkeys() string {
	s := ""

	for _, v := range l.payload {
		s = s + string(v.key)
	}

	return s
}

func (l *level) getvalues() interface{} {
	s := []interface{}{}

	for _, v := range l.payload {
		s = append(s, v.value)
	}

	return s
}

// newLevel
func newLevel() *level {
	return &level{
		payload: []cell{},
	}
}

// Dict whatever
type Dict interface {
	Set(string, interface{})
	Get(string) interface{}
}

func (d *dict) Set(key string, v interface{}) {

}

func (d *dict) Get(key string) interface{} {
	return nil
}

func (d *dict) insert(key string, value interface{}) {

}

// New constructor
func New() Dict {
	d := &dict{}
	d.root = d
	return d
}

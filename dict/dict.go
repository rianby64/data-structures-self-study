package dict

type dict struct {
	payload interface{}
	key     string
	root    *dict
	left    *dict
	right   *dict
	parent  *dict
}

type level struct {
	changed bool
	min     byte
	max     byte
	len     int
	payload []byte
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
	if actual == b {
		return index, true
	}

	if actual < b {
		for actual != b {
			index++
			if index >= l.len {
				return -1, false
			}
			actual = l.payload[index]
		}
		return index, true
	}

	for actual != b {
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
	l.min = l.payload[0]
	l.max = l.payload[l.len-1]
	l.changed = false
}

func (l *level) insert(b byte) {
	defer l.updateEdges()

	if b > l.max {
		l.changed = true
		l.payload = append(l.payload, b)
		return
	}

	if b < l.min {
		l.changed = true
		l.payload = append([]byte{b}, l.payload...)
		return
	}

	if b == l.max || b == l.min {
		return
	}

	newpayload := []byte{}
	added := false

	for _, v := range l.payload {
		if v == b {
			return
		}

		if v > b && !added {
			newpayload = append(newpayload, b)
			added = true
		}
		newpayload = append(newpayload, v)
	}

	l.changed = true
	l.payload = newpayload
}

func (l *level) delete(b byte) {
	defer l.updateEdges()

	newpayload := []byte{}
	changed := false

	for i := 0; i < l.len; i++ {
		v := l.payload[i]
		if v == b {
			changed = true
			continue
		}

		newpayload = append(newpayload, v)
	}

	l.changed = changed
	l.payload = newpayload
}

func (l *level) String() string {
	return string(l.payload)
}

// newLevel
func newLevel() *level {
	return &level{
		payload: []byte{},
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

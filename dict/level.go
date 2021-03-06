package dict

const (
	half float64 = 2
)

type cell struct {
	key   byte
	value interface{}
	child *level
}

type level struct {
	changed bool
	min     byte
	max     byte
	len     int
	payload []*cell
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
	iq := (float64(1) - (q / d))

	i := float64(l.len) * (ip + iq) / half
	index := int(i)

	actual := l.payload[index]

	if actual.key == b {
		return index, true
	}

	if actual.key < b {
		for actual.key < b {
			index++
			actual = l.payload[index]
		}
	}

	if actual.key > b {
		for actual.key > b {
			index--
			actual = l.payload[index]
		}
	}

	if actual.key == b {
		return index, true
	}

	return -1, false
}

func (l *level) updateEdges() {
	if !l.changed {
		return
	}

	l.len = len(l.payload)

	if l.len == 0 {
		l.min = 0
		l.max = 0
	} else {
		l.min = l.payload[0].key
		l.max = l.payload[l.len-1].key
	}

	l.changed = false
}

func (l *level) insert(key byte, value interface{}) (c *cell) {
	defer func() {
		if c.value == nil || value != nil {
			c.value = value
		}

		l.updateEdges()
	}()

	if key == l.min {
		c = l.payload[0]

		return c
	}

	if key == l.max {
		c = l.payload[l.len-1]

		return c
	}

	c = &cell{key: key}
	l.changed = true

	if key > l.max {
		l.payload = append(l.payload, c)

		return c
	}

	if key < l.min {
		l.payload = append([]*cell{c}, l.payload...)

		return c
	}

	newpayload := []*cell{}
	added := false

	for i, v := range l.payload {
		if v.key == key {
			c = l.payload[i]
			l.changed = false

			return c
		}

		if v.key > key && !added {
			c = &cell{key: key}
			newpayload = append(newpayload, c)
			added = true
		}

		newpayload = append(newpayload, v)
	}

	l.payload = newpayload

	return c
}

func (l *level) delete(key byte) bool {
	if key < l.min || key > l.max {
		return false
	}

	defer l.updateEdges()

	newpayload := []*cell{}
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

	return changed
}

func (l *level) getkeys() string {
	s := ""

	for _, v := range l.payload {
		s += string(v.key)
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

func (l *level) updatevalue(key byte, value interface{}) bool {
	i, ok := l.getIndex(key)
	if !ok {
		return false
	}

	l.payload[i].value = value

	return true
}

func (l *level) getvalue(key byte) (interface{}, bool) {
	i, ok := l.getIndex(key)
	if !ok {
		return nil, false
	}

	return l.payload[i].value, true
}

// newLevel constructor.
func newLevel() *level {
	return &level{
		payload: []*cell{},
	}
}

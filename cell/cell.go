package cell

// Cell interface
type Cell interface {
	Value() interface{}
	SetValue(v interface{})
}

type cell struct {
	payload interface{}
}

func (c *cell) Value() interface{} {
	return c.payload
}

func (c *cell) SetValue(v interface{}) {
	c.payload = v
}

// New constructor
func New(payload interface{}) Cell {
	return &cell{
		payload: payload,
	}
}

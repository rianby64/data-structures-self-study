package dict

type dict struct {
	payload interface{}
	key     string
	root    *dict
	left    *dict
	right   *dict
	parent  *dict
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

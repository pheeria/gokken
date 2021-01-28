package gokken

type Dict struct {
	keys   []string
	values []string
	size   int
}

func (d *Dict) hash(key string) int {
	return len(key)
}
func (d *Dict) Init() {
	d.keys = make([]string, d.size, d.size*2)
	d.values = make([]string, d.size, d.size*2)
}
func (d *Dict) Set(key, value string) {
	hashed := d.hash(key)
	d.keys[hashed] = key
	d.values[hashed] = value
}
func (d *Dict) Get(key string) string {
	hashed := d.hash(key)
	return d.values[hashed]
}

func (d Dict) String() string {
	result := "{\n"

	for i, v := range d.keys {
		if v != "" {
			result += "  " + v + ": " + d.values[i] + "\n"
		}
	}
	result += "}"

	return result
}

package circular_int

type Circle struct {
	elements_vector []int
	contor          int
}

func NewCircle() *Circle {
	return &Circle{}
}

func (c *Circle) Next() int {
	if c.contor == len(c.elements_vector) {
		c.contor = 0
	}
	return c.elements_vector[c.contor]
}
func (c *Circle) Len() int {
	return len(c.elements_vector)
}
func (c *Circle) Append(data int) {
	c.elements_vector = append(c.elements_vector, data)
}
func (c *Circle) GerElements() []int {
	return c.elements_vector
}

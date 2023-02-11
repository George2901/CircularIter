package circular_byte

type Circle struct {
	elements_vector []byte
	contor          int
}

func NewCircle() *Circle {
	return &Circle{}
}

func (c *Circle) Next() byte {
	if c.contor == len(c.elements_vector) {
		c.contor = 0
	}
	return c.elements_vector[c.contor]
}
func (c *Circle) Len() int {
	return len(c.elements_vector)
}
func (c *Circle) Append(data byte) {
	c.elements_vector = append(c.elements_vector, data)
}
func (c *Circle) GerElements() []byte {
	return c.elements_vector
}

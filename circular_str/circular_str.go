package circular_str

type Circle struct {
	elements_vector []string
	contor          int
}

func NewCircle() *Circle {
	return &Circle{}
}

func (c *Circle) Next() string {
	if c.contor == len(c.elements_vector) {
		c.contor = 0
	}
	return c.elements_vector[c.contor]
}
func (c *Circle) Len() int {
	return len(c.elements_vector)
}
func (c *Circle) Append(data string) {
	c.elements_vector = append(c.elements_vector, data)
}
func (c *Circle) GerElements() []string {
	return c.elements_vector
}

package circle

type CircleInt struct {
	elements_vector []int
	contor          int
}

func NewCircleInt() *CircleInt {
	return &CircleInt{}
}

func (c *CircleInt) Next() int {
	if c.contor == len(c.elements_vector) {
		c.contor = 0
	}
	return c.elements_vector[c.contor]
}
func (c *CircleInt) Len() int {
	return len(c.elements_vector)
}
func (c *CircleInt) Append(data int) {
	c.elements_vector = append(c.elements_vector, data)
}
func (c *CircleInt) GerElements() []int {
	return c.elements_vector
}

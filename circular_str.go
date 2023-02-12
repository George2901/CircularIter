package circle

type CircleStr struct {
	elements_vector []string
	contor          int
}

func NewCircleStr() *CircleStr {
	return &CircleStr{}
}

func (c *CircleStr) Next() string {
	if c.contor == len(c.elements_vector) {
		c.contor = 0
	}
	return c.elements_vector[c.contor]
}
func (c *CircleStr) Len() int {
	return len(c.elements_vector)
}
func (c *CircleStr) Append(data string) {
	c.elements_vector = append(c.elements_vector, data)
}
func (c *CircleStr) GerElements() []string {
	return c.elements_vector
}

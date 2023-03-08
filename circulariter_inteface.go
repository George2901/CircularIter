package circle

type CircleInterface struct {
	elements_vector []interface{}
	contor          int
}

func NewCircleInterface() *CircleInterface {

	return &CircleInterface{}
}

func (c *CircleInterface) Next() interface{} {
	if c.contor == len(c.elements_vector) {
		c.contor = 0
	}
	return c.elements_vector[c.contor]
}
func (c *CircleInterface) Len() int {
	return len(c.elements_vector)
}
func (c *CircleInterface) Append(data interface{}) {
	c.elements_vector = append(c.elements_vector, data)
}
func (c *CircleInterface) GerElements() []interface{} {
	return c.elements_vector
}

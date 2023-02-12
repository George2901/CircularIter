package circle

type CircleByte struct {
	elements_vector []byte
	contor          int
}

func NewCircleByte() *CircleByte {
	return &CircleByte{}
}

func (c *CircleByte) Next() byte {
	if c.contor == len(c.elements_vector) {
		c.contor = 0
	}
	return c.elements_vector[c.contor]
}
func (c *CircleByte) Len() int {
	return len(c.elements_vector)
}
func (c *CircleByte) Append(data byte) {
	c.elements_vector = append(c.elements_vector, data)
}
func (c *CircleByte) GerElements() []byte {
	return c.elements_vector
}

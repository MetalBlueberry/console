package console

func (c *Console) Assert(assertion bool, obj ...any) {
	if assertion {
		c.Error(obj...)
	}
}

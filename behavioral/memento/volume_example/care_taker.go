package main

type careTaker struct {
	mementoStack []memento
}

func (c *careTaker) Add(m memento) {
	c.mementoStack = append(c.mementoStack, m)
}

func (c *careTaker) Pop() memento {
	if len(c.mementoStack) > 0 {
		tempMemento := c.mementoStack[len(c.mementoStack)-1]
		c.mementoStack = c.mementoStack[0 : len(c.mementoStack)-1]
		return tempMemento
	}
	return memento{}
}

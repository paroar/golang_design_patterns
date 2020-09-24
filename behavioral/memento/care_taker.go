package memento

import "errors"

type careTaker struct {
	mementoList []memento
}

func (c *careTaker) Add(m memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) Memento(i int) (memento, error) {
	if len(c.mementoList) < i || i < 0 {
		return memento{}, errors.New("Index not found\n")
	}
	return c.mementoList[i], nil
}

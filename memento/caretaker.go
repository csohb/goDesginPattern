package main

type CareTaker struct {
	mementoArray []*Memento
}

func (c *CareTaker) addMemento(m *Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *CareTaker) getMemento(index int) *Memento {
	return c.mementoArray[index]
}

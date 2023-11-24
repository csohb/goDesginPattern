package main

type Originator struct {
	speed int
}

func (o *Originator) getSpeed() int {
	return o.speed
}

func (o *Originator) setSpeed(speed int) {
	o.speed = speed
}

func (o *Originator) createMemento() *Memento {
	return &Memento{speed: o.speed}
}

func (o *Originator) restoreMemento(m *Memento) {
	o.speed = m.getSavedSpeed()
}

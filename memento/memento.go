package main

type Memento struct {
	speed int
}

func (m *Memento) getSavedSpeed() int {
	return m.speed
}

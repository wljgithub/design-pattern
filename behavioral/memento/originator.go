package main

type  originator struct {
	state string
}

func (this *originator)createMemento() *memento  {
	return &memento{state: this.state}
}

func (this *originator)restoreMemento(m *memento)  {
	this.state = m.getSavedState()
}

func (this *originator) setState(state string) {
	this.state = state
}

func (this *originator) getState() string {
	return this.state
}

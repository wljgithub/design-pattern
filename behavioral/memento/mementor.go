package main

type memento struct {
	state string
}

func (this *memento)getSavedState() string  {
	return this.state
}

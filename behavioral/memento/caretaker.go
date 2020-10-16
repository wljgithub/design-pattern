package main

type caretaker struct {
	mementoArray []*memento
}

func (this *caretaker) addMemento(m *memento)  {
	this.mementoArray = append(this.mementoArray,m)
}

func (this *caretaker) getMemento(index int) *memento {
	if index <0 || index >= len(this.mementoArray){
		return nil
	}
	return this.mementoArray[index]
}

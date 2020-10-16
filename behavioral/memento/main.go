package main

import "fmt"

func main() {
	caretaker := &caretaker{mementoArray: make([]*memento,0)}

	originator := &originator{state: "A"}
	fmt.Printf("Originator current state: %s\n",originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("B")
	fmt.Printf("Originator current state: %s\n",originator.getState())

	caretaker.addMemento(originator.createMemento())
	originator.setState("C")

	fmt.Printf("Originator current state: %s\n",originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.getMemento(1))
	fmt.Printf("Restored to state: %s\n",originator.getState())

	originator.restoreMemento(caretaker.getMemento(0))
	fmt.Printf("Restore to state: %s\n",originator.getState())
}




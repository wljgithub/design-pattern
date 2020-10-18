package main

import "fmt"

type item struct {
	observerList []observer
	name string
	inStock bool
}

func newItem(name string) *item {
	return &item{name: name}
}

func (this *item)updateAvailability()  {
	fmt.Printf("Item %s is now in stock\n",this.name)
	this.inStock = true
	this.notifyAll()
}

func (this *item) register(o observer) {
	this.observerList = append(this.observerList,o)
}

func (this *item) deregister(o observer) {
	this.observerList = removeFromSlice(this.observerList,o)
}

func (this *item)notifyAll()  {
	for _,observer := range this.observerList{
		observer.update(this.name)
	}
}

func removeFromSlice(observerList []observer, observerToRemove observer) []observer {
	sliceLen := len(observerList)
	for i,observer := range observerList{
		if observer.getID() == observerToRemove.getID(){
			observerList[i],observerList[sliceLen-1] = observerList[sliceLen-1],observerList[i]
			return observerList[:sliceLen-1]
		}
	}
	return observerList
}
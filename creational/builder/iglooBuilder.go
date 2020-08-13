package main

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (this *iglooBuilder) setWindowType() {
	this.windowType = "Snow Door"
}

func (this *iglooBuilder) setNumFloor() {
	this.floor = 1
}

func (this *iglooBuilder) setDoorType() {
	this.doorType = "Snow Door"
}

func (this *iglooBuilder) getHouse() house {
	return house{
		doorType:   this.doorType,
		windowType: this.windowType,
		floor:      this.floor,
	}
}

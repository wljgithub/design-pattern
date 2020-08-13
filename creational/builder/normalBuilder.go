package main

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (this *normalBuilder) setWindowType() {
	this.windowType = "Wooden Window"
}

func (this *normalBuilder) setDoorType() {
	this.doorType = "Wooden Door"
}

func (this *normalBuilder) setNumFloor() {
	this.floor = 2
}

func (this *normalBuilder) getHouse() house {
	return house{
		doorType:   this.doorType,
		windowType: this.windowType,
		floor:      this.floor,
	}
}

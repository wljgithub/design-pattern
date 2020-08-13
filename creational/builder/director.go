package main

type director struct {
	builder iBuilder
}

func newDirector(this iBuilder) *director {
	return &director{
		builder: this,
	}
}

func (this *director) setBuilder(b iBuilder) {
	this.builder = b
}

func (this *director) buildHouse() house {
	this.builder.setDoorType()
	this.builder.setWindowType()
	this.builder.setNumFloor()
	return this.builder.getHouse()
}

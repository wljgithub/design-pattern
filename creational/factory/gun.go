package main

type gun struct {
	name  string
	power int
}

func (this *gun) setName(name string) {
	this.name = name
}

func (this *gun) setPower(power int) {
	this.power = power
}

func (this *gun) getName() string {
	return this.name
}

func (this *gun) getPower() int {
	return this.power
}

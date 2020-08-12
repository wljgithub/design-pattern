package main

type iShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}
type shoe struct {
	logo string
	size int
}

func (this *shoe) setLogo(logo string) {
	this.logo = logo
}

func (this *shoe) setSize(size int) {
	this.size = size
}

func (this *shoe) getLogo() string {
	return this.logo
}

func (this *shoe) getSize() int {
	return this.size
}

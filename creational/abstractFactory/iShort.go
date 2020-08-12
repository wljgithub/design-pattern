package main

type iShort interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type short struct {
	logo string
	size int
}

func (this *short) setLogo(logo string) {
	this.logo = logo
}

func (this *short) setSize(size int) {
	this.size = size
}

func (this *short) getLogo() string {
	return this.logo
}

func (this *short) getSize() int {
	return this.size
}

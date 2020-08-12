package main

type adidas struct {
}

func (this *adidas) makeShoe() iShoe {
	shoe := &adidasShoe{}
	shoe.setLogo("adidas")
	shoe.setSize(24)
	return shoe
}

func (this *adidas) makeShort() iShort {
	short := &adidasShort{}
	short.setLogo("adidas")
	short.setSize(24)
	return short
}

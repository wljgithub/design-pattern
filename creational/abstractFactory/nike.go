package main

type nike struct {
}

func (n *nike) makeShoe() iShoe {
	shoe := &nikeShoe{}
	shoe.setLogo("nike")
	shoe.setSize(24)
	return shoe
}

func (n *nike) makeShort() iShort {
	short := &nikeShort{}
	short.setLogo("nike")
	short.setSize(24)
	return short
}

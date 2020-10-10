package main

type couterTerroristDress struct {
	color string
}

func (c *couterTerroristDress) getColor() string {
	return c.color
}

func newCouterTerroristDress() *couterTerroristDress {
	return &couterTerroristDress{color:"green"}
}
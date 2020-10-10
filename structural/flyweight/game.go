package main

type game struct {
	players []*player
}

func newGame() *game {
	return &game{}
}

func (g *game) addTerrorist(dressType string) {
	player := newPlayer("Terrorist",dressType)
	g.players = append(g.players,player)
}

func (g *game) addCouterTerrorist(dressType string) {
	player := newPlayer("CouterTerrorist",dressType)
	g.players = append(g.players,player)
}

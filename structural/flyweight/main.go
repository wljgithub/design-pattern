package main

import "fmt"

func main() {
	game := newGame()


	// add terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	// add counter-terrorist
	game.addCouterTerrorist(CounterTerroristDressType)
	game.addCouterTerrorist(CounterTerroristDressType)
	game.addCouterTerrorist(CounterTerroristDressType)
	dressFactorySingleInstance := getDressFactorySingleInstance()
	for dressType ,dress := range dressFactorySingleInstance.dressMap{
		fmt.Printf("DressColorType:%s DressColor:%s\n",dressType,dress)
	}
}

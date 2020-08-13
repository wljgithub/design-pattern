package main

import "fmt"

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Println("Normal House Door Type:",normalHouse.doorType)
	fmt.Println("Normal House Window Type:",normalHouse.windowType)
	fmt.Println("Normal House Num Floor:",normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Println("Igloo House Door Type:",iglooHouse.doorType)
	fmt.Println("Igloo House Windos Type:",iglooHouse.windowType)
	fmt.Println("Igloo House Num Floor:",iglooHouse.floor)
}

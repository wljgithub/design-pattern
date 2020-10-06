package main

import "fmt"

func main() {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}
	macComputer := &mac{}
	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	fmt.Println()

	winComputer := &windows{}
	winComputer.setPrinter(hpPrinter)
	winComputer.print()
	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
	fmt.Println()
}

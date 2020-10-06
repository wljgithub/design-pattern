package main

import "fmt"

type epson struct {

}

func (p *epson) printFile() {
	fmt.Println("printing by a EPSON Printer")
}

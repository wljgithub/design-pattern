package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	mavarick, _ := getGun("mavarick")

	printDetails(ak47)
	printDetails(mavarick)
}

func printDetails(gun iGun) {
	fmt.Println("Gun:", gun.getName())
	fmt.Println("Power:", gun.getPower())
}

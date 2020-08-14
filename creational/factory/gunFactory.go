package main

import "fmt"

func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAK47(), nil
	}
	if gunType == "mavarick" {
		return newMaverick(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

package main

import "fmt"

type iSportsFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "nike" {
		return &nike{}, nil
	}

	if brand == "adidas" {
		return &adidas{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}

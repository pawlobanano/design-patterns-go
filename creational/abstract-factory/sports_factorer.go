package main

import "fmt"

type SportsFactorer interface {
	makeShoe() Shoer
	makeShirt() Shirter
}

func sportsFactory(brand string) (SportsFactorer, error) {
	switch brand {
	case "adadis":
		return &Adadis{}, nil
	case "neki":
		return &Neki{}, nil
	default:
		return &Dummy{}, fmt.Errorf("wrong brand type provided")
	}
}

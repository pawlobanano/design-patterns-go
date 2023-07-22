package main

import "fmt"

func getGun(gunType string) (Guner, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "glock":
		return newGlock(), nil
	default:
		return &Gun{}, fmt.Errorf("wrong gun type provided")
	}
}

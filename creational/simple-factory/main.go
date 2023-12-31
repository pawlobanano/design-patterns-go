package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	glock, _ := getGun("glock")

	printDetails(ak47)
	printDetails(glock)
}

func printDetails(g Gunner) {
	fmt.Printf("Gun: %s\n", g.Name())
	fmt.Printf("Power: %d\n", g.Power())
}

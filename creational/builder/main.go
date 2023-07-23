package main

import "fmt"

func main() {
	normalHouseBuilder := getHouseBuilder("normal")
	constructionManager := newConstructionManager(normalHouseBuilder)
	normalHouse := constructionManager.buildHouse()

	fmt.Printf("Normal house door type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal house window type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal house floors: %d\n\n", normalHouse.floors)

	iglooHouseBuilder := getHouseBuilder("igloo")
	constructionManager.setHouseBuilder(iglooHouseBuilder)
	iglooHouse := constructionManager.buildHouse()

	fmt.Printf("Igloo house door type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo house window type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo house floors: %d\n\n", iglooHouse.floors)

	dummyHouseBuilder := getHouseBuilder("dummy")
	constructionManager.setHouseBuilder(dummyHouseBuilder)
	dummyHouse := constructionManager.buildHouse()

	fmt.Printf("Dummy house door type: %s\n", dummyHouse.doorType)
	fmt.Printf("Dummy house window type: %s\n", dummyHouse.windowType)
	fmt.Printf("Dummy house floors: %d\n", dummyHouse.floors)
}

package main

import "fmt"

func main() {
	normalHouseBuilder := getHouseBuilder("normall")

	constructionManager := newConsturtionManager(normalHouseBuilder)
	normalHouse := constructionManager.buildHouse()

	fmt.Printf("Normal house door type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal house window type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal house floors: %d\n", normalHouse.floors)
}

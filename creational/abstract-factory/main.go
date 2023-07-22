package main

import "fmt"

func main() {
	adadisFactory, _ := sportsFactory("adadis")
	nekiFactory, _ := sportsFactory("neki")
	dummyFactory, _ := sportsFactory("non-existent-brand")

	adadisShoe := adadisFactory.makeShoe()
	adadisShirt := adadisFactory.makeShirt()
	nekiShoe := nekiFactory.makeShoe()
	nekiShirt := nekiFactory.makeShirt()
	dummyShoe := dummyFactory.makeShoe()
	dummyShirt := dummyFactory.makeShirt()

	printShoeDetails(adadisShoe)
	printShirtDetails(adadisShirt)
	printShoeDetails(nekiShoe)
	printShirtDetails(nekiShirt)
	printShoeDetails(dummyShoe)
	printShirtDetails(dummyShirt)
}

func printShoeDetails(s Shoer) {
	fmt.Printf("Shoe logo: %s\n", s.logo())
	fmt.Printf("Shoe size: %d\n\n", s.size())
}

func printShirtDetails(s Shirter) {
	fmt.Printf("Shirt logo: %s\n", s.logo())
	fmt.Printf("Shirt size: %d\n\n", s.size())
}

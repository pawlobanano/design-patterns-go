package main

import "fmt"

func main() {
	adadisFactory, _ := SportsFactory("adadis")
	nekiFactory, _ := SportsFactory("neki")

	adadisShoe := adadisFactory.makeShoe()
	adadisShirt := adadisFactory.makeShirt()
	nekiShoe := nekiFactory.makeShoe()
	nekiShirt := nekiFactory.makeShirt()

	printShoeDetails(adadisShoe)
	printShirtDetails(adadisShirt)
	printShoeDetails(nekiShoe)
	printShirtDetails(nekiShirt)
}

func printShoeDetails(s Shoer) {
	fmt.Printf("Shoe logo: %s\n", s.logo())
	fmt.Printf("Shoe size: %d\n\n", s.size())
}

func printShirtDetails(s Shirter) {
	fmt.Printf("Shirt logo: %s\n", s.logo())
	fmt.Printf("Shirt size: %d\n\n", s.size())
}

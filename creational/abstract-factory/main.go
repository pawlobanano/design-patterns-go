package main

import "fmt"

func main() {
	adadisFactory, _ := SportsFactory("adadis")

	adadisShoe := adadisFactory.makeShoe()
	adadisShirt := adadisFactory.makeShirt()

	printShoeDetails(adadisShoe)
	printShirtDetails(adadisShirt)
}

func printShoeDetails(s Shoer) {
	fmt.Printf("Shoe logo: %s\n", s.logo())
	fmt.Printf("Shoe size: %d\n", s.size())
}

func printShirtDetails(s Shirter) {
	fmt.Printf("Shirt logo: %s\n", s.logo())
	fmt.Printf("Shirt size: %d\n", s.size())
}

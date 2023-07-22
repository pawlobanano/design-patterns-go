package main

type Dummy struct{}

func (a *Dummy) makeShoe() Shoer {
	return &DummyShoe{}
}

func (a *Dummy) makeShirt() Shirter {
	return &DummyShirt{}
}

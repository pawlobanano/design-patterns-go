package main

type Adadis struct{}

func (a *Adadis) makeShoe() Shoer {
	return &AdadisShoe{
		Shoe: Shoe{
			logo_: "adadis",
			size_: 10,
		},
	}
}

func (a *Adadis) makeShirt() Shirter {
	return &AdadisShirt{
		Shirt: Shirt{
			logo_: "adadis",
			size_: 50,
		},
	}
}

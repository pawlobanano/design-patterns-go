package main

type Neki struct{}

func (a *Neki) makeShoe() Shoer {
	return &NekiShoe{
		Shoe: Shoe{
			logo_: "neki",
			size_: 8,
		},
	}
}

func (a *Neki) makeShirt() Shirter {
	return &NekiShirt{
		Shirt: Shirt{
			logo_: "neki",
			size_: 38,
		}}
}

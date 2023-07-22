package main

type glock struct {
	Gun
}

func newGlock() Guner {
	return &glock{
		Gun: Gun{
			name:  "Glock pistol",
			power: 1,
		},
	}
}

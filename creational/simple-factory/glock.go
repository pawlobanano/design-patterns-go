package main

type glock struct {
	Gun
}

func newGlock() Gunner {
	return &glock{
		Gun: Gun{
			name:  "Glock pistol",
			power: 1,
		},
	}
}

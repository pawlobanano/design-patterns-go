package main

type glock struct {
	Gun
}

func newGlock() IGun {
	return &glock{
		Gun: Gun{
			name:  "Glock pistol",
			power: 1,
		},
	}
}

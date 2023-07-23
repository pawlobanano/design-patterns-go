package main

type ak47 struct {
	Gun
}

func newAk47() Gunner {
	return &ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

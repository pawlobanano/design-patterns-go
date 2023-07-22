package main

type ak47 struct {
	Gun
}

func newAk47() Guner {
	return &ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

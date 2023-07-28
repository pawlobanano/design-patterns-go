package main

type NormalHouseBuilder struct {
	windowType string
	doorType   string
	floors     int
}

func newNormalHouseBuilder() *NormalHouseBuilder {
	return &NormalHouseBuilder{}
}

func (b *NormalHouseBuilder) setWindowType() {
	b.windowType = "Wooden window"
}

func (b *NormalHouseBuilder) setDoorType() {
	b.doorType = "Wooden door"
}

func (b *NormalHouseBuilder) setFloors() {
	b.floors = 2
}

func (b *NormalHouseBuilder) build() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floors:     b.floors,
	}
}

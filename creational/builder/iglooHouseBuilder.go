package main

type IglooHouseBuilder struct {
	windowType string
	doorType   string
	floors     int
}

func newIglooHouseBuilder() *IglooHouseBuilder {
	return &IglooHouseBuilder{}
}

func (b *IglooHouseBuilder) setWindowType() {
	b.windowType = "Wooden window"
}

func (b *IglooHouseBuilder) setDoorType() {
	b.doorType = "Wooden door"
}

func (b *IglooHouseBuilder) setFloors() {
	b.floors = 2
}

func (b *IglooHouseBuilder) build() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floors:     b.floors,
	}
}

package main

type Dummy struct {
	house House
}

func (d *Dummy) setWindowType() {
	d.house.windowType = ""
}

func (d *Dummy) setDoorType() {
	d.house.doorType = ""
}

func (d *Dummy) setFloors() {
	d.house.floors = 0
}

func (a *Dummy) build() House {
	return House{}
}

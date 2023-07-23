package main

type Builderer interface {
	setWindowType()
	setDoorType()
	setFloors()
	build() House
}

func getHouseBuilder(builderType string) Builderer {
	switch builderType {
	case "normal":
		return newNormalHouseBuilder()
	case "igloo":
		return newIglooHouseBuilder()
	default:
		return &Dummy{}
	}
}

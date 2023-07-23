package main

type ConstructionManager struct {
	houseBuilder Builderer
}

func newConstructionManager(b Builderer) *ConstructionManager {
	return &ConstructionManager{
		houseBuilder: b,
	}
}

func (cM *ConstructionManager) setHouseBuilder(b Builderer) {
	cM.houseBuilder = b
}

func (cM *ConstructionManager) buildHouse() House {
	cM.houseBuilder.setDoorType()
	cM.houseBuilder.setWindowType()
	cM.houseBuilder.setFloors()

	return cM.houseBuilder.build()
}

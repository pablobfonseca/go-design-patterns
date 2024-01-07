package main

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return &NormalBuilder{}
	}

	if builderType == "igloo" {
		return &IglooBuilder{}
	}

	return nil
}

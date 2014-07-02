package gocamp

import "testing"

func factoryWorld() *World {
	w := &World{}
	w.createWorld(3, 3, 1)
	return w
}

func TestCreateWorld(t *testing.T) {
	w := &World{}
	w.createWorld(3, 3, 1)
}

func TestShowPlane(t *testing.T) {
	w := factoryWorld()
	w.showTruePlane(1)
}

func TestWestOffset(t *testing.T) {
	w := factoryWorld()
	// I don't know how to test these yet, perhaps I need a type representing a
	// point on the map
}

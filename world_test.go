package gocamp

import "testing"

func TestCreateWorld(t *testing.T) {
	w := &World{}
	w.createWorld(3, 3, 1)
}

func TestShowPlane(t *testing.T) {
	w := &World{}
	w.createWorld(3, 3, 1)
	w.showTruePlane(1)
}

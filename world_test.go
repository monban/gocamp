package gocamp

import "testing"

func factoryWorld() World {
	w := createWorld(3, 6, 1)
	return w
}

func TestCreateWorld(t *testing.T) {
	w := createWorld(3, 3, 1)
	if &w == nil {
		t.Errorf("expected w to be a world")
	}
}
/*
func TestShowPlane(t *testing.T) {
	w := factoryWorld()
	w.showTruePlane(1)
}

func TestWestOffset(t *testing.T) {
	w := factoryWorld()
	w.showTruePlane(1)
	// I don't know how to test these yet, perhaps I need a type representing a
	// point on the map
}
*/
func TestPtToIndex(t *testing.T) {
	 // create a world with a true xyz size of 3x3x3
	w := createWorld(1,1,1)
	// create a point at the one valid space
	point := Pt{0,0,0}
	if w.ptToIndex(point) != 4 {
		t.Errorf("expected 4 got %d", w.ptToIndex(point))
	}
}

func TestSetBlock(t *testing.T) {
	w := factoryWorld()
	blok := Block{rock,rock}
	pt := Pt{1,1,1}
	w.setBlock(pt, blok)
	if w.getBlock(pt) != blok {
		t.Errorf("expected %s got %s", blok, w.getBlock(pt))
	}
}

func TestFillPlane(t *testing.T) {
	w:= factoryWorld()
	blok := Block{rock,rock}
	w.fillPlane(0, blok)
	if w.getBlock(Pt{1,1,0}) != blok {
		t.Errorf("expected %s got %s", blok, w.getBlock(Pt{1,1,0}))
	}
}

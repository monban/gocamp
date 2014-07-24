package gocamp

import "testing"

func factoryTerrain() Terrain {
	w := createTerrain(3, 6, 1)
	return w
}

func TestCreateTerrain(t *testing.T) {
	w := createTerrain(3, 3, 1)
	if &w == nil {
		t.Errorf("expected w to be a world")
	}
}

/*
func TestShowPlane(t *testing.T) {
	w := factoryTerrain()
	w.showTruePlane(1)
}

func TestWestOffset(t *testing.T) {
	w := factoryTerrain()
	w.showTruePlane(1)
	// I don't know how to test these yet, perhaps I need a type representing a
	// point on the map
}
*/
func TestPtToIndex(t *testing.T) {
	// create a world with a true xyz size of 3x3x3
	w := createTerrain(1, 1, 1)
	// create a point at the one valid space
	point := Pt{0, 0, 0}
	if w.ptToIndex(point) != 4 {
		t.Errorf("expected 4 got %d", w.ptToIndex(point))
	}
}

func TestSetBlock(t *testing.T) {
	w := factoryTerrain()
	blok := Block{rock, rock}
	pt := Pt{1, 1, 1}
	w.setBlock(pt, blok)
	if w.getBlock(pt) != blok {
		t.Errorf("expected %s got %s", blok, w.getBlock(pt))
	}
}

func TestFillPlane(t *testing.T) {
	w := factoryTerrain()
	blok := Block{rock, rock}
	w.fillPlane(0, blok)
	if w.getBlock(Pt{1, 1, 0}) != blok {
		t.Errorf("expected %s got %s", blok, w.getBlock(Pt{1, 1, 0}))
	}
}

func TestFillAndShow(t *testing.T) {
	w := factoryTerrain()
	w.fillPlane(1, Block{air, air})
	// TODO: Assertion?
	//w.showTruePlane(1)
}

func TestBlocksToRunes(t *testing.T) {
	//	w := factoryTerrain()
	//plane, _ := w.getPlane(1)
	//	_ = blocksToRunes(plane)
	//TODO: assertion
}

func TestLevelAsRuneArray(t *testing.T) {
	w := factoryTerrain()
	ra := w.LevelAsRuneArray(1)

	// Test the size of the outer array
	if len(ra) != w.trueSizeY() {
		t.Errorf("expected outer array to be length %d, but was %d", w.trueSizeY(), len(ra))
	}

	// Test the size of the inner array
	if len(ra[0]) != w.trueSizeX() {
		t.Errorf("expected inner array to be length %d, but was %d", w.trueSizeX(), len(ra[0]))
	}

	// TODO: Test the content of the arrays
}

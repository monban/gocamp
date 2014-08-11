package gocamp

import "testing"

func FactoryWorld() *defaultWorld {
	return createDefaultWorld(4, 4, 1)
}

func TestCreateWorld(t *testing.T) {
	w := FactoryWorld()
	if &w == nil {
		t.Errorf("expected w to be a world")
	}
}

func TestEntitiesThink(t *testing.T) {
	w := FactoryWorld()
	e := createStaticEntity()
	ei := Entitier(&e)
	w.AddEntity(ei)
	cycles := 32
	for i := 0; i < cycles; i++ {
		w.Tick()
	}

	if e.cycles != cycles {
		t.Errorf("expected cycles to be %d but got %d", cycles, e.cycles)
	}
}

func TestGetEntireLevelAsRuneArray(t *testing.T) {
	w := FactoryWorld()
	e := CreateStaticEntity()
	w.AddEntity(e)
	ra := w.GetEntireLevelAsRuneArray(1)
	if ra[e.GetPosition().x][e.GetPosition().y] != e.DisplayRune() {
		t.Errorf("entity should be displaying rune %v at %d,%d, found %v instead",
			e.DisplayRune(), e.GetPosition().x, e.GetPosition().y, ra[e.GetPosition().x][e.GetPosition().y])
	}
}

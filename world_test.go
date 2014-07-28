package gocamp

import "testing"

func FactoryWorld() *defaultWorld {
	w := new(defaultWorld)
	w.terrain = createTerrain(4, 4, 1)
	w.size_x = 4
	w.size_y = 4
	w.size_z = 1
	return w
}

func TestCreateWorld(t *testing.T) {
	w := FactoryWorld()
	if &w == nil {
		t.Errorf("expected w to be a world")
	}
}

func TestEntitiesThink(t *testing.T) {
	w := FactoryWorld()
	e := staticEntity{}
	ei := Entitier(&e)
	w.entities.Add(&ei)
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
	w.entities.Add(&e)
	ra := w.GetEntireLevelAsRuneArray(1)
	if ra[e.GetPosition().x][e.GetPosition().y] != e.DisplayRune() {
		t.Errorf("entity should be displaying rune %v at %d,%d, found %v instead",
			e.DisplayRune(), e.GetPosition().x, e.GetPosition().y, ra[e.GetPosition().x][e.GetPosition().y])
	}
}

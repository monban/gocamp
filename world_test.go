package gocamp

import "testing"

func FactoryWorld() *world {
	return CreateWorld(4, 4, 1)
}

func TestCreateWorld(t *testing.T) {
	w := FactoryWorld()
	if &w == nil {
		t.Errorf("expected w to be a world")
	}
}

func TestEntitiesThink(t *testing.T) {
	w := FactoryWorld()
	e := new(StaticEntity)
	w.entities.Add(e)
	cycles := 32
	for i := 0; i < cycles; i++ {
		w.Tick()
	}

	if e.cycles != cycles {
		t.Errorf("expected cycles to be %d but got %d", cycles, e.cycles)
	}
}

package gocamp

import "testing"

func FactoryWorld() *world {
	return CreateWorld(4,4,1)
}

func TestCreateWorld(t *testing.T) {
	w := FactoryWorld()
	if &w == nil {
		t.Errorf("expected w to be a world")
	}
}

func TestAddEntities(t *testing.T) {
	w := FactoryWorld()
	e := new(StaticEntity)
	w.AddEntity(e)
	if !w.EntityExists(e) {
		t.Errorf("entity did not get added to list")
	}
}

func TestEntitiesThink(t *testing.T) {
	w := FactoryWorld()
	e := new(StaticEntity)
	w.AddEntity(e)
	
	cycles := 32
	for i := 0; i < cycles; i++ {
		w.Tick()
	}
	
	if e.cycles != cycles {
		t.Errorf("expected cycles to be %d but got %d", cycles, e.cycles)
	}
}

func TestMakeLotsOfEntities(t *testing.T) {
	const number int = 32
	w := FactoryWorld()
	for i := 0; i < number; i++ {
		w.AddEntity(new(StaticEntity))
	}
	if len(w.entities) == number {
		t.Errorf("expected %d entities but got %d", number, len(w.entities))
	}
}

package gocamp
import "testing"

func TestAddEntities(t *testing.T) {
	em := new(EntityManager)
	e := new(StaticEntity)
	em.Add(e)
	if !em.Exists(e) {
		t.Errorf("entity did not get added to list")
	}
	if em.Exists(new(StaticEntity)) {
		t.Errorf("newly created entity somehow ended up in the list")
	}
}

func TestMakeLotsOfEntities(t *testing.T) {
	const number int = 32
	em := new(EntityManager)
	for i := 0; i < number; i++ {
		em.Add(new(StaticEntity))
	}
	if len(em.entities) != number {
		t.Errorf("expected %d entities but got %d", number, len(em.entities))
	}
}

func TestGetRune(t *testing.T) {
	e := CreateTestStaticEntity()
	if e.DisplayRune() != e.displayRune {
		t.Errorf("expected %d got %d", e.displayRune, e.DisplayRune())
	}
}

func TestWalker(t *testing.T) {
	e := CreateWalkerEntity(Pt{1,1,1}, e)
	for i := 1; i < 10; i++ {
		if (e.location != Pt{i,1,1}) {
			t.Errorf("expected %d got %d", e.location.x, i)
		}
		e.Think()
	}
}

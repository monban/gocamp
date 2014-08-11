package gocamp

import "testing"

func TestAddEntities(t *testing.T) {
	em := CreateDefaultEntityManager()
	e := CreateStaticEntity()
	em.Add(e)
	if !em.Exists(e) {
		t.Errorf("entity did not get added to list")
	}

	new_se := CreateStaticEntity()
	if em.Exists(new_se) {
		t.Errorf("newly created entity somehow ended up in the list")
	}
}

func TestMakeLotsOfEntities(t *testing.T) {
	const number int = 32
	em := CreateDefaultEntityManager()
	for i := 0; i < number; i++ {
		e := CreateStaticEntity()
		em.Add(e)
	}
	if len(em.All()) != number {
		t.Errorf("expected %d entities but got %d", number, len(em.All()))
	}
}

func TestGetRune(t *testing.T) {
	e := staticEntity{}
	if e.DisplayRune() != e.displayRune {
		t.Errorf("expected %d got %d", e.displayRune, e.DisplayRune())
	}
}

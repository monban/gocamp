package gocamp

import "testing"

func TestWalker(t *testing.T) {
	w := createDefaultWorld(6, 6, 1)
	em := w.GetEntityManager()
	walker := createWalkerEntity()
	walker.SetDestination(Pt{2, 2, 1})
	e := Entitier(&walker)
	(*em).Add(&e)
	for i := 0; i < 3; i++ {
		w.Tick()
		PrintRuneArray(w.GetEntireLevelAsRuneArray(1))
	}
}

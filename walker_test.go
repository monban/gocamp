package gocamp

import "testing"

func TestWalker(t *testing.T) {
	w := createDefaultWorld(6, 6, 1)
	em := w.EntityManager()
	walker := createWalkerEntity()
	walker.SetDestination(Pt{2, 2, 1})
	em.Add(&walker)
	walker.SetWorld(w) // HACK BUG FIXME: WE SHOULD NOT HAVE TO DO THIS
	for i := 0; i < 3; i++ {
		w.Tick()
		PrintRuneArray(w.GetEntireLevelAsRuneArray(1))
	}
}

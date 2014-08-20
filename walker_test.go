package gocamp

import "testing"

func TestWalker(t *testing.T) {
	w := createDefaultWorld(6, 6, 1)
	w.terrain.fillPlane(1, Block{air, rock})
	w.terrain.NewLine(Block{rock, rock}, Pt{0, 1, 1}, 5, true)
	w.terrain.NewLine(Block{rock, rock}, Pt{1, 3, 1}, 5, true)
	walker := createWalkerEntity()
	walker.SetDestination(Pt{6, 6, 1})
	w.AddEntity(&walker)
	for walker.location != walker.destination {
		w.Tick()
		PrintRuneArray(w.GetEntireLevelAsRuneArray(1))
	}
}

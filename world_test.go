package gocamp

import "testing"

func TestCreateWorld(t *testing.T) {
	w := CreateWorld(4,4,1)
	if &w == nil {
		t.Errorf("expected w to be a world")
	}
}

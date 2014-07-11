package gocamp

import "errors"
import "fmt"

type material int

const (
	void material = iota + 1 // The void is complete emptiness
	air                      // Air is empty space, things live here
	soil                     // Dirt
	rock                     // What it says on the tin
)

type dir int
const (
	w dir = iota
	e
	n
	s
	u
	d
	nw
	ne
	sw
	se
	unw
	une
	usw
	use
	dnw
	dne
	dsw
	dse
)

type Block struct {
	bulk  material
	floor material
}

func (b *Block) ascii_representation() string {
	switch b.bulk {
	case void:
		return "*"
	case air:
		return " "
	case soil:
		return "."
	case rock:
		return "x"
	}
	panic("invalid material found in bulk")
}

type World struct {
	size_x      int
	size_y    int
	size_z      int
	blocks     []Block   // list of all blocks
}

type Pt struct {
	x int
	y int
	z int
}

func (w *World) trueWidth() int {
	return w.size_x + 2
}

func (w *World) trueBreadth() int {
	return w.size_y + 2
}

func (w *World) trueDepth() int {
	return w.size_z + 2
}

func (self *World) dirOffset(direction dir) int {
	switch direction {
	case w:
		return -1
	case e:
		return 1
	case n:
		return -(self.trueWidth())
	case s:
		return -(self.trueWidth())
	case u:
		return -(self.trueWidth() * self.trueBreadth())
	case d:
		return self.trueWidth() * self.trueBreadth()
	}
	return 0
}

func (w *World) createWorld(size_x int, size_y int, size_z int) {
	w.size_x = size_x
	w.size_y = size_y
	w.size_z = size_z

	numberBlocks := w.trueWidth() * w.trueBreadth() * w.trueDepth()

	w.blocks = make([]Block, numberBlocks)

	for i := 0; i < len(w.blocks); i++ {
		w.blocks[i] = Block{void, void}
	}
}

// Returns a slice containing the world at a certain z-level
func (w *World) getPlane(z_level int) (level []Block, err error) {
	if z_level < 1 || z_level > w.size_z {
		return nil, errors.New("no such z level")
	}
	begin := w.trueWidth() * w.trueDepth()
	return w.blocks[begin : begin+w.dirOffset('d')], nil
}

// Print out everything on a certain level, mostly for debugging
func (w *World) showTruePlane(z_level int) {
	plane, _ := w.getPlane(z_level)
	for x := 0; x < w.trueWidth(); x++ {
		for y := 0; y < w.trueBreadth(); y++ {
			fmt.Printf(plane[(y*w.trueWidth())+x].ascii_representation())
		}
		fmt.Printf("\n")
	}
}

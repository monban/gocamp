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
	width      int
	breadth    int
	depth      int
	blocks     []Block   // list of all blocks
	diagOffset [12]Block // distance to each diagonal direction
}

func (w *World) trueWidth() int {
	return w.width + 2
}

func (w *World) trueBreadth() int {
	return w.breadth + 2
}

func (w *World) trueDepth() int {
	return w.depth + 2
}

func (w *World) dirOffset(dir rune) int {
	switch dir {
	case 'w':
		return -1
	case 'e':
		return 1
	case 'n':
		return -(w.trueWidth())
	case 's':
		return -(w.trueWidth())
	case 'u':
		return -(w.trueWidth() * w.trueBreadth())
	case 'd':
		return w.trueWidth() * w.trueBreadth()
	}
	return 0
}

func (w *World) createWorld(size_x int, size_y int, size_z int) {
	w.width = size_x
	w.breadth = size_y
	w.depth = size_z

	numberBlocks := w.trueWidth() * w.trueBreadth() * w.trueDepth()

	w.blocks = make([]Block, numberBlocks)

	for i := 0; i < len(w.blocks); i++ {
		w.blocks[i] = Block{void, void}
	}
}

// Returns a slice containing the world at a certain z-level
func (w *World) getPlane(z_level int) (level []Block, err error) {
	if z_level < 1 || z_level > w.depth {
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

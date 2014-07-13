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

func (w *World) trueSizeX() int {
	return w.size_x + 2
}

func (w *World) trueSizeY() int {
	return w.size_y + 2
}

func (w *World) trueSizeZ() int {
	return w.size_z + 2
}

func (self *World) dirOffset(direction dir) int {
	switch direction {
	case w:
		return -1
	case e:
		return 1
	case n:
		return -(self.trueSizeX())
	case s:
		return -(self.trueSizeX())
	case u:
		return -(self.trueSizeX() * self.trueSizeY())
	case d:
		return self.trueSizeX() * self.trueSizeY()
	}
	return 0
}

func createWorld(size_x int, size_y int, size_z int) World {
	w := new(World)
	w.size_x = size_x
	w.size_y = size_y
	w.size_z = size_z

	numberBlocks := w.trueSizeX() * w.trueSizeY() * w.trueSizeZ()

	w.blocks = make([]Block, numberBlocks)

	for i := 0; i < len(w.blocks); i++ {
		w.blocks[i] = Block{void, void}
	}
	return *w
}

// Returns a slice containing the world at a certain z-level
func (w *World) getPlane(z_level int) (level []Block, err error) {
	if z_level < 1 || z_level > w.size_z {
		return nil, errors.New("no such z level")
	}
	begin := w.trueSizeX() * w.trueSizeZ()
	return w.blocks[begin : begin+w.dirOffset(d)], nil
}

// Print out everything on a certain level, mostly for debugging
func (w *World) showTruePlane(z_level int) {
	plane, _ := w.getPlane(z_level)
	for x := 0; x < w.trueSizeX(); x++ {
		for y := 0; y < w.trueSizeY(); y++ {
			fmt.Printf(plane[(y*w.trueSizeX())+x].ascii_representation())
		}
		fmt.Printf("\n")
	}
}

// Converts point coordinates to an array index
func (self *World) ptToIndex(point Pt) int {
	// Skip the top border
	i := self.trueSizeX()

	// Skip the left padding and move to the correct column
	i += point.x+1

	// Move to the correct row
	i += point.y*self.trueSizeX()
	
	// Move to the correct level
	i += point.z*self.trueSizeX()*self.trueSizeY()
	
	return i
}

func (self *World) getBlock(pt Pt) Block {
	return self.blocks[self.ptToIndex(pt)]
}

func (self *World) setBlock(pt Pt, blok Block) {
	self.blocks[self.ptToIndex(pt)] = blok
}

func (self *World) fillPlane(z_level int, bulk material) {
	for x := 0; x < self.size_x; x++ {
		for y := 0; y < self.size_y; y++ {
			self.setBlock(Pt{x,y,z_level}, Block{bulk, bulk})
		}
	}
}

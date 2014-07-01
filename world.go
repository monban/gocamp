package gocamp

import "errors"
import "fmt"

type material int

const (
  void material = iota + 1    // The void is complete emptiness
  air                         // Air is empty space, things live here
  soil                        // Dirt
  rock                        // What it says on the tin
)

type Block struct {
  bulk    material
  floor   material
}

func (b* Block) ascii_representation() (string) {
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
  width       int
  breadth     int
  depth       int
  trueWidth   int
  trueBreadth int
  trueDepth   int
  blocks      []Block     // list of all blocks
  dirOffset   [6]int      // distance to each cardinal direction
  diagOffset  [12]Block   // distance to each diagonal direction
}


func (w* World) createWorld(size_x int, size_y int, size_z int) {
  w.width = size_x
  w.breadth = size_y
  w.depth = size_z

  w.trueWidth = w.width + 2
  w.trueBreadth = w.breadth + 2
  w.trueDepth = w.depth + 2

  numberBlocks := w.trueWidth * w.trueBreadth * w.trueDepth
  
  // Setup offsets
  w.dirOffset[0] = 1              //w
  w.dirOffset[1] = -1             //e
  w.dirOffset[2] = w.trueWidth    //s
  w.dirOffset[3] = -w.trueWidth   //n
  w.dirOffset[4] = w.trueWidth * w.trueBreadth
  w.dirOffset[5] = -(w.trueWidth * w.trueBreadth)
  
  w.blocks = make([]Block, numberBlocks)
  
  for i := 0; i< len(w.blocks); i++ {
    w.blocks[i] = Block{void, void}
  }
}

// Returns a slice containing the world at a certain z-level
func (w* World) getPlane(z_level int) (level []Block, err error) {
  if z_level < 1 || z_level > w.depth {
    return nil, errors.New("no such z level")
  }
  begin := w.trueWidth * w.trueDepth
  return w.blocks[begin:begin+w.dirOffset[4]], nil
}

 // Print out everything on a certain level, mostly for debugging
func (w* World) showTruePlane(z_level int) {
  plane,_ := w.getPlane(z_level)
  for x:=0; x < w.trueWidth; x++ {
    for y:=0; y < w.trueBreadth; y++ {
      fmt.Printf(plane[(y*w.trueWidth)+x].ascii_representation())
    }
    fmt.Printf("\n")
  }
}

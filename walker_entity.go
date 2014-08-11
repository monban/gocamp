package gocamp

import "github.com/sperre/astar"

func CreateWalkerEntity() Entitier {
	we := createWalkerEntity()
	return Entitier(&we)
}

func createWalkerEntity() WalkerEntity {
	we := new(WalkerEntity)
	we.location = Pt{1, 1, 1}
	we.destination = Pt{1, 1, 1}
	return *we
}

type WalkerEntity struct {
	location    Pt
	cycles      int
	destination Pt
	path        []Pt
	world       WorldRepresenter
}

func (self *WalkerEntity) GetPosition() Pt {
	return self.location
}

func (self *WalkerEntity) Think() {
	if self.location == self.destination {
		return
	}
	if len(self.path) == 0 {
		self.pathfind()
	}
	self.moveAlongPath()
}

func (self *WalkerEntity) DisplayRune() rune {
	return '@'
}

func (self *WalkerEntity) SetWorld(world WorldRepresenter) {
	self.world = world
}

func (self *WalkerEntity) SetDestination(destination Pt) {
	self.destination = destination
}

func (self *WalkerEntity) pathfind() {
	map_data := self.world.GetTraversableMap(self.location.z)
	path := astar.Astar(map_data, self.location.x, self.location.y, self.destination.x, self.destination.y, false)
	self.path = make([]Pt, len(path))
	for i, node := range path {
		self.path[i] = Pt{node.X, node.Y, self.location.y}
	}
}

func (self *WalkerEntity) moveAlongPath() {
	if len(self.path) == 0 {
		return
	}
	self.location = self.path[0]
	self.path = self.path[1:len(self.path)]
}

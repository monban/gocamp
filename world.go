package gocamp

func CreateWorld(size_x int, size_y int, size_z int) *world {
	w := new(world)
	w.terrain = createTerrain(size_x, size_y, size_z)
	return w
}

type world struct {
	size_x int
	size_y int
	size_z int
	terrain Terrain
	entities []Entitier
}

// TODO: This function will eventually use goroutines to sync entities thinking
// (or perhaps be run in a goroutine itself)
func (self *world) Tick() {
	for i, _ :=  range self.entities {
		self.entities[i].Think()
	}
}

func (self *world) AddEntity(entity Entitier) {
	self.entities = append(self.entities, entity)
}

func (self *world) EntityExists(entity Entitier) bool {
	for _, foo := range self.entities {
		if entity == foo {
			return true
		}
	}
	return false
}

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
	entities EntityManager
}

// TODO: This function will eventually use goroutines to sync entities thinking
// (or perhaps be run in a goroutine itself)
func (self *world) Tick() {
	self.entities.Tick()
}

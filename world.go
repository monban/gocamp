package gocamp

func CreateWorld(size_x int, size_y int, size_z int) *world {
	w := new(world)
	w.terrain = createTerrain(size_x, size_y, size_z)
	w.size_x = size_x
	w.size_y = size_y
	w.size_z = size_z
	return w
}

type world struct {
	size_x   int
	size_y   int
	size_z   int
	terrain  Terrain
	entities EntityManager
}

// TODO: This function will eventually use goroutines to sync entities thinking
// (or perhaps be run in a goroutine itself)
func (self *world) Tick() {
	self.entities.Tick()
}

// Return a level as an array of runes
// This is mostly a development / debug function
func (self *world) GetEntireLevelAsRuneArray(level int) [][]rune {
	output := self.terrain.LevelAsRuneArray(level)
	for i := range self.entities.entities {
		e := self.entities.entities[i]
		if e.GetPosition().z != level {
			continue
		}
		output[e.GetPosition().x][e.GetPosition().y] = e.DisplayRune()
	}
	return output
}

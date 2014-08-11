package gocamp

func CreateDefaultWorld(size_x int, size_y int, size_z int) WorldRepresenter {
	w := createDefaultWorld(size_x, size_y, size_z)
	return WorldRepresenter(w)
}

func createDefaultWorld(size_x int, size_y int, size_z int) *defaultWorld {
	w := new(defaultWorld)
	terrain := createTerrain(size_x, size_y, size_z)
	w.terrain = &terrain
	w.size_x = size_x
	w.size_y = size_y
	w.size_z = size_z
	w.entities = CreateDefaultEntityManager()
	return w
}

type defaultWorld struct {
	size_x   int
	size_y   int
	size_z   int
	terrain  *Terrain
	entities EntityManager
}

// TODO: This function will eventually use goroutines to sync entities thinking
// (or perhaps be run in a goroutine itself)
func (self *defaultWorld) Tick() {
	self.entities.Tick()
}

// Return a level as an array of runes
// This is mostly a development / debug function
func (self *defaultWorld) GetEntireLevelAsRuneArray(level int) [][]rune {
	output := self.terrain.LevelAsRuneArray(level)
	for i := range self.entities.All() {
		e := self.entities.All()[i]
		if e.GetPosition().z != level {
			continue
		}
		output[e.GetPosition().x][e.GetPosition().y] = e.DisplayRune()
	}
	return output
}

func (self *defaultWorld) AddEntity(entity Entitier) {
	self.entities.Add(entity)
	wr := WorldRepresenter(self)
	entity.SetWorld(wr)
}

func (self *defaultWorld) EntityManager() EntityManager {
	return self.entities
}

func (self *defaultWorld) GetTerrain() *Terrain {
	return self.terrain
}

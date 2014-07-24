package gocamp

type EntityManager struct {
	entities []Entitier
}

func (self *EntityManager) Add(entity Entitier) {
	self.entities = append(self.entities, entity)
}

func (self *EntityManager) Exists(entity Entitier) bool {
	for _, foo := range self.entities {
		if entity == foo {
			return true
		}
	}
	return false
}

func (self *EntityManager) Tick() {
	for i, _ :=  range self.entities {
		self.entities[i].Think()
	}
}


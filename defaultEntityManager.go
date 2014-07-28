package gocamp

func CreateDefaultEntityManager() EntityManager {
	em := defaultEntityManager{}
	return EntityManager(&em)
}

type defaultEntityManager struct {
	entities []Entitier
}

func (self *defaultEntityManager) Add(entity *Entitier) {
	self.entities = append(self.entities, *entity)
}

func (self *defaultEntityManager) Exists(entity Entitier) bool {
	for _, foo := range self.entities {
		if entity == foo {
			return true
		}
	}
	return false
}

func (self *defaultEntityManager) Tick() {
	for i, _ := range self.entities {
		self.entities[i].Think()
	}
}

func (self *defaultEntityManager) All() []Entitier {
	return self.entities
}

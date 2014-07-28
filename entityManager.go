package gocamp

type EntityManager interface {
	Add(entity *Entitier)
	Exists(entity *Entitier) bool
	Tick()
	All() []*Entitier
}

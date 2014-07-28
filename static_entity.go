package gocamp

func CreateStaticEntity() Entitier {
	se := createStaticEntity()
	return Entitier(&se)
}

func createStaticEntity() staticEntity {
	se := staticEntity{}
	se.location = Pt{1, 1, 1}
	se.cycles = 0
	se.displayRune = '☺'
	return se
}

type staticEntity struct {
	location    Pt
	cycles      int
	displayRune rune
}

func (self *staticEntity) GetPosition() Pt {
	return self.location
}

func (self *staticEntity) Think() {
	self.cycles++
}

func (self *staticEntity) DisplayRune() rune {
	return self.displayRune
}

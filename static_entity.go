package gocamp

func CreateTestStaticEntity() StaticEntity {
	se := new(StaticEntity)
	se.location = Pt{1,1,1}
	se.cycles = 0
	se.displayRune = '☺'
	return *se
}

type StaticEntity struct {
	location Pt
	cycles int
	displayRune rune
}

func (self *StaticEntity) GetPosition() Pt {
	return self.location
}

func (self *StaticEntity) Think() {
	self.cycles++
}

func (self *StaticEntity) DisplayRune() rune {
	return self.displayRune
}

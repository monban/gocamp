package gocamp

func CreateTestStaticEntity() StaticEntity {
	se := new(StaticEntity)
	se.location = Pt{1,1,1}
	se.cycles = 0
	return *se
}

type StaticEntity struct {
	location Pt
	cycles int
}

func (self *StaticEntity) GetPosition() Pt {
	return self.location
}

func (self *StaticEntity) Think() {
	self.cycles++
}

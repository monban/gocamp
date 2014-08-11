package gocamp

func CreateWalkerEntity() Entitier {
	we := createWalkerEntity()
	return Entitier(&we)
}

func createWalkerEntity() WalkerEntity {
	we := new(WalkerEntity)
	we.location = Pt{1, 1, 1}
	we.destination = Pt{1, 1, 1}
	return *we
}

type WalkerEntity struct {
	location    Pt
	cycles      int
	destination Pt
}

func (self *WalkerEntity) GetPosition() Pt {
	return self.location
}

func (self *WalkerEntity) Think() {

}

func (self *WalkerEntity) DisplayRune() rune {
	return '☺'
}

func (self *WalkerEntity) SetWorld(world *WorldRepresenter) {
}

func (self *WalkerEntity) SetDestination(destination Pt) {
	self.destination = destination
}

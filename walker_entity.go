package gocamp

func CreateWalkerEntity(start Pt, direction dir) WalkerEntity {
	we := new(WalkerEntity)
	we.location = start
	we.direction = direction
	return *we
}

type WalkerEntity struct {
	location Pt
	cycles int
	direction dir
}

func (self *WalkerEntity) GetPosition() Pt {
	return self.location
}

func (self *WalkerEntity) Think() {
	self.location.Move(self.direction)
}

func (self *WalkerEntity) DisplayRune() rune {
	return '☺'
}

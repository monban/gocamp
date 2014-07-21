package gocamp

// This is just a terrible name for this interface don't you think?
type Entitier interface {
	// All Entities must keep track of their own position in the world
	// GetPosition() returns a Pt containing the current location of the entity
	GetPosition() Pt

	// Every tick this function is called on every entity
	Think()

	// Represent yourself as a rune
	// TODO: Need to make this more interface independant
	DisplayRune() rune
}

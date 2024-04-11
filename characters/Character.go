package characters

import (
	"image/color"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// The character class is a super class for both the player and the NPC class
type Character struct {
	Position mazegrid.Position
	Count    int
	Colour   color.Color
}

// Initialises the character given the grid position and a colour
func (c *Character) Init(startPos mazegrid.Position, colour color.Color) {

	c.Position = startPos
	c.Colour = colour

}

// Function that returns the current position of the character
func (c *Character) GetPosition() mazegrid.Position {
	return c.Position
}

// Function that sets a new position for the character
func (c *Character) SetPosition(p mazegrid.Position) {
	c.Position = p
}

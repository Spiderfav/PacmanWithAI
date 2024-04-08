package characters

import (
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// The character class is a super class for both the player and the NPC class
type Character struct {
	Sprite   *ebiten.Image
	Position mazegrid.Position
	Count    int
	Colour   color.Color
}

// Initialises the character given the grid position and a colour
func (c *Character) Init(startPos mazegrid.Position, colour color.Color) {

	c.Sprite = ebiten.NewImageFromImage(c.setSprite())
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

// Function that updates the counter for the sprite
func (c *Character) UpdateCount() {
	c.Count += 1
}

// Function that returns the counter for the sprite of the character
func (c *Character) GetCount() int {
	return c.Count
}

// Function that returns the sprite of the character
func (c *Character) GetSprite() *ebiten.Image {
	return c.Sprite
}

// Functions that sets the sprite of the character
func (c *Character) setSprite() image.Image {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

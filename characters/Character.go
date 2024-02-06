package characters

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

type FrameProperties struct {
	FrameOX     int
	FrameOY     int
	FrameWidth  int
	FrameHeight int
	FrameCount  int
}

type Character struct {
	Sprite   *ebiten.Image
	Position mazegrid.Position
	Algo     int
	FrameProperties
	Count int
}

func (c *Character) Init(pos mazegrid.Position) {

	c.Sprite = ebiten.NewImageFromImage(c.setSprite())
	c.Position = pos
	c.Algo = 0
	c.FrameOX = 0
	c.FrameOY = 32
	c.FrameWidth = 32
	c.FrameHeight = 32
	c.FrameCount = 8
}

func (c *Character) GetAlgo() int {
	return c.Algo
}

func (c *Character) GetPosition() mazegrid.Position {
	return c.Position
}

func (c *Character) SetPosition(p mazegrid.Position) {
	c.Position = p
}

func (c *Character) GetFrameProperties() FrameProperties {
	return c.FrameProperties
}

func (c *Character) SetFrameProperties(fp FrameProperties) {
	c.FrameProperties = fp
}

func (c *Character) UpdateCount() {
	c.Count += 1
}

func (c *Character) GetCount() int {
	return c.Count
}

func (c *Character) GetSprite() *ebiten.Image {
	return c.Sprite
}

func (c *Character) setSprite() image.Image {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

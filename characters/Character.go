package characters

import (
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	FrameProperties
	Count  int
	Colour color.Color
}

func (c *Character) Init(startPos mazegrid.Position, colour color.Color) {

	c.Sprite = ebiten.NewImageFromImage(c.setSprite())
	c.Position = startPos
	c.FrameOX = 0
	c.FrameOY = 32
	c.FrameWidth = 32
	c.FrameHeight = 32
	c.FrameCount = 8
	c.Colour = colour

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

func DrawSprite(screen *ebiten.Image, char Character) {
	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(float64(char.GetPosition().XCoordinate+10), float64((char.GetPosition().YCoordinate + 10)))
	// i := (char.GetCount() / 5) % char.GetFrameProperties().FrameCount
	// sx, sy := char.GetFrameProperties().FrameOX+i*char.GetFrameProperties().FrameWidth, char.GetFrameProperties().FrameOY
	// screen.DrawImage(char.GetSprite().SubImage(image.Rect(sx, sy, sx+char.GetFrameProperties().FrameWidth, sy+char.GetFrameProperties().FrameHeight)).(*ebiten.Image), op)

	vector.DrawFilledCircle(screen, char.GetPosition().XCoordinate+10, char.GetPosition().YCoordinate+10, 2, char.Colour, true)

}

package characters

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// This method means type NPC implements the interface Character,
// but I don't need to explicitly declare that it does so.
type NPC struct {
	Sprite      *ebiten.Image
	Position    mazegrid.Position
	Algo        int
	frameOX     int
	frameOY     int
	frameWidth  int
	frameHeight int
	frameCount  int
}

func (n NPC) GetPosition() mazegrid.Position {
	return n.Position

}

func (n NPC) SetPosition(newPos mazegrid.Position) {
	n.Position = newPos
}

func (n NPC) SetSprite() {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}

	n.Sprite = ebiten.NewImageFromImage(img)
}

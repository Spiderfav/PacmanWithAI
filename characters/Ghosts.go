package characters

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// This method means type NPC implements the interface Character,
// but I don't need to explicitly declare that it does so.
type NPC struct {
	Atributes CharacterAtributes
}

func CreateGhost() NPC {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}

	n := NPC{CharacterAtributes{ebiten.NewImageFromImage(img), mazegrid.Position{}, 0, 0, 32, 32, 32, 8}}
	return n

}

func (n NPC) GetPosition() mazegrid.Position {
	return n.Atributes.Position

}

func (n NPC) SetPosition(newPos mazegrid.Position) {
	n.Atributes.Position = newPos
}

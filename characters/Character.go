package characters

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

type Character interface {
	GetPosition() mazegrid.Position
	SetPosition(mazegrid.Position)
	SetSprite()
}

type CharacterAtributes struct {
	Sprite      *ebiten.Image
	Position    mazegrid.Position
	Algo        int
	FrameOX     int
	FrameOY     int
	FrameWidth  int
	FrameHeight int
	FrameCount  int
}

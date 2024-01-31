package characters

import (
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

type Character interface {
	GetPosition() mazegrid.Position
	SetPosition(mazegrid.Position)
	SetSprite()
}

package characters

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

type DirectionOfPlayer = int

const (
	Up    DirectionOfPlayer = 0
	Down  DirectionOfPlayer = 1
	Left  DirectionOfPlayer = 2
	Right DirectionOfPlayer = 3
)

type Player struct {
	Attributes Character
	Points     int
}

func (p *Player) Init(startPos mazegrid.Position, colour color.Color) {
	p.Attributes.Init(startPos, colour)

}

func (p *Player) GetPosition() mazegrid.Position {
	return p.Attributes.GetPosition()
}

func (p *Player) SetPosition(m mazegrid.Position) {
	p.Attributes.SetPosition(m)
}

func (p *Player) GetFrameProperties() FrameProperties {
	return p.Attributes.GetFrameProperties()
}

func (p *Player) SetFrameProperties(fp FrameProperties) {
	p.Attributes.SetFrameProperties(fp)
}

func (p *Player) UpdateCount() {
	p.Attributes.UpdateCount()
}

func (p *Player) GetCount() int {
	return p.Attributes.GetCount()
}

func (p *Player) GetSprite() *ebiten.Image {
	return p.Attributes.GetSprite()
}

func (p *Player) ResetPoints() {
	p.Points = 0
}

func (p *Player) GetPoints() int {
	return p.Points
}

func (p *Player) Move(d DirectionOfPlayer, m [][]mazegrid.MazeSquare, squareSize int) {

	array2Pos := int((p.Attributes.Position.XCoordinate / float32(squareSize)) - 1)
	array1Pos := int((p.Attributes.Position.YCoordinate / float32(squareSize)) - 1)

	if m[array1Pos][array2Pos].HasPellot {
		p.Points += 1
		m[array1Pos][array2Pos].HasPellot = false
	}

	if m[array1Pos][array2Pos].HasSuperPellot {
		p.Points += 5
		m[array1Pos][array2Pos].HasSuperPellot = false
	}

	switch d {
	case Up:
		if !m[array1Pos][array2Pos].HasWalls.HasUp {
			p.Attributes.SetPosition(m[array1Pos][array2Pos].Walls.Up)
		}

	case Down:
		if !m[array1Pos][array2Pos].HasWalls.HasDown {
			p.Attributes.SetPosition(m[array1Pos][array2Pos].Walls.Down)
		}

	case Right:
		if !m[array1Pos][array2Pos].HasWalls.HasRight {
			p.Attributes.SetPosition(m[array1Pos][array2Pos].Walls.Right)
		}

	case Left:
		if !m[array1Pos][array2Pos].HasWalls.HasLeft {
			p.Attributes.SetPosition(m[array1Pos][array2Pos].Walls.Left)
		}
	}

}

// func (p *Player) setSprite() image.Image {
// 	return p.Attributes.setSprite()
// }

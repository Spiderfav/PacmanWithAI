package characters

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

type DirectionOfPlayer = int

// These are the directions available for a player to take
const (
	Up    DirectionOfPlayer = 0
	Down  DirectionOfPlayer = 1
	Left  DirectionOfPlayer = 2
	Right DirectionOfPlayer = 3
)

// This player object is an extension of character
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

// This function, given the direction that the player is moving, will move the player one square in that direction
func (p *Player) move(d DirectionOfPlayer, m [][]mazegrid.MazeSquare, squareSize int) {

	array2Pos := int((p.Attributes.Position.XCoordinate / float32(squareSize)) - 1)
	array1Pos := int((p.Attributes.Position.YCoordinate / float32(squareSize)) - 1)

	// If the square contains a pellot, add points to the player
	// Delete the pellot from the square once collected
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

// This function checks if the player is moving and moves the player in the game grid if so.
func (p *Player) IsPlayerMoving(gameGrid [][]mazegrid.MazeSquare, squareSize int) {

	// Checking if the player is moving and if so, moving the player
	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		p.move(Up, gameGrid, squareSize)

	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {

		// for !inpututil.IsKeyJustReleased(ebiten.KeyS) || inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) {
		p.move(Down, gameGrid, squareSize)
		//}

	} else if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		p.move(Left, gameGrid, squareSize)

	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		p.move(Right, gameGrid, squareSize)

	}

}

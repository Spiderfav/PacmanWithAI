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
	Attributes  Character
	mapPoints   int
	totalPoints int
	lives       int
}

// Function to initialise a player. Requires the starting position, colour and lives for the player.
func (p *Player) Init(startPos mazegrid.Position, colour color.Color, lives int) {
	p.Attributes.Init(startPos, colour)
	p.lives = lives

}

// Function to return position of the player
func (p *Player) GetPosition() mazegrid.Position {
	return p.Attributes.GetPosition()
}

// Function to set position of the player
func (p *Player) SetPosition(m mazegrid.Position) {
	p.Attributes.SetPosition(m)
}

// Function that resets the total points of the player
func (p *Player) addAllPoints(pointsToAdd int) {
	p.totalPoints += pointsToAdd
	p.mapPoints += pointsToAdd
}

// Function that resets the total points of the player
func (p *Player) ResetAllPoints() {
	p.totalPoints = 0
	p.mapPoints = 0
}

// Function that resets the total points of the player
func (p *Player) ResetTotalPoints() {
	p.totalPoints = 0
}

// Function that returns the current total points of the player
func (p *Player) GetTotalPoints() int {
	return p.totalPoints
}

// Function that resets the points of the player for that map
func (p *Player) ResetMapPoints() {
	p.mapPoints = 0
}

// Function that returns the current points of the player for the current map
func (p *Player) GetMapPoints() int {
	return p.mapPoints
}

// Function that resets the lives of the player
func (p *Player) ResetLives() {
	p.lives = 3
}

// Function that returns the current lives of the player
func (p *Player) GetLives() int {
	return p.lives
}

// Function that removes a life from the player
func (p *Player) RemoveLife() {
	p.lives -= 1
}

// Function that calls all necessary functions on a game over
func (p *Player) GameOver() {
	p.ResetLives()
	p.ResetAllPoints()
}

// This function, given the direction that the player is moving, will move the player one square in that direction
func (p *Player) move(d DirectionOfPlayer, m [][]mazegrid.MazeSquare, squareSize int) {

	array2Pos := int((p.Attributes.Position.XCoordinate / float32(squareSize)) - 1)
	array1Pos := int((p.Attributes.Position.YCoordinate / float32(squareSize)) - 1)

	// If the square contains a pellot, add points to the player
	// Delete the pellot from the square once collected
	if m[array1Pos][array2Pos].HasPellot {
		p.addAllPoints(1)
		m[array1Pos][array2Pos].HasPellot = false
	}

	if m[array1Pos][array2Pos].HasSuperPellot {
		p.addAllPoints(5)
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

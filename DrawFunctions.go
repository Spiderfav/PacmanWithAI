package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/characters"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// This function draws a given square to the screen
// It checks if the current node has a given wall, then draws it to the screen
func drawSquare(screen *ebiten.Image, squareToDraw mazegrid.MazeSquare) {
	var strokeWidth float32 = 1

	if squareToDraw.ContainsObject {
		vector.DrawFilledCircle(screen, squareToDraw.NodePosition.XCoordinate+float32(halfSquare), squareToDraw.NodePosition.YCoordinate+float32(halfSquare), halfSquare/4, color.RGBA{255, 100, 0, 250}, true)
	}

	if squareToDraw.HasWalls.HasDown {
		vector.StrokeLine(screen, squareToDraw.NodePosition.XCoordinate, squareToDraw.NodePosition.YCoordinate+float32(squareSize), squareToDraw.NodePosition.XCoordinate+float32(squareSize), squareToDraw.NodePosition.YCoordinate+float32(squareSize), strokeWidth, color.Black, false)
	}

	if squareToDraw.HasWalls.HasRight {
		vector.StrokeLine(screen, squareToDraw.NodePosition.XCoordinate+float32(squareSize), squareToDraw.NodePosition.YCoordinate, squareToDraw.NodePosition.XCoordinate+float32(squareSize), squareToDraw.NodePosition.YCoordinate+float32(squareSize), strokeWidth, color.Black, false)
	}

	if squareToDraw.HasWalls.HasLeft {
		vector.StrokeLine(screen, squareToDraw.NodePosition.XCoordinate, squareToDraw.NodePosition.YCoordinate, squareToDraw.NodePosition.XCoordinate, squareToDraw.NodePosition.YCoordinate+float32(squareSize), strokeWidth, color.Black, false)
	}

	if squareToDraw.HasWalls.HasUp {
		vector.StrokeLine(screen, squareToDraw.NodePosition.XCoordinate, squareToDraw.NodePosition.YCoordinate, squareToDraw.NodePosition.XCoordinate+float32(squareSize), squareToDraw.NodePosition.YCoordinate, strokeWidth, color.Black, false)
	}

}

// The DrawMaze function takes the screen argument given as the screen to draw to maze to
// It draws the maze from the GameGridDFS
func drawMaze(screen *ebiten.Image, g *Game) {

	// For each row and column, it looks at the walls of the block and draws the ones it has
	for i := 0; i < g.Maze.Size; i++ {
		for j := 0; j < g.Maze.Size; j++ {
			drawSquare(screen, g.Maze.Grid[i][j])
		}
	}
}

// This function draws lines to the screen for a given Ghost/Pacman path
func drawPathsLines(screen *ebiten.Image, pathTaken []mazegrid.MazeSquare) {
	if len(pathTaken) == 0 {
		return
	}
	prevX := pathTaken[0].NodePosition.XCoordinate + halfSquare
	prevY := pathTaken[0].NodePosition.YCoordinate + halfSquare

	for i := 1; i < len(pathTaken); i++ {
		vector.StrokeLine(screen, prevX, prevY, pathTaken[i].NodePosition.XCoordinate+halfSquare, pathTaken[i].NodePosition.YCoordinate+halfSquare, 1, color.RGBA{0, 255, 0, 250}, false)
		prevX = pathTaken[i].NodePosition.XCoordinate + halfSquare
		prevY = pathTaken[i].NodePosition.YCoordinate + halfSquare

	}

}

// Function to draw the sprites of the characters to the screen
func DrawSprite(screen *ebiten.Image, char characters.Character) {
	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(float64(char.GetPosition().XCoordinate+10), float64((char.GetPosition().YCoordinate + 10)))
	// i := (char.GetCount() / 5) % char.GetFrameProperties().FrameCount
	// sx, sy := char.GetFrameProperties().FrameOX+i*char.GetFrameProperties().FrameWidth, char.GetFrameProperties().FrameOY
	// screen.DrawImage(char.GetSprite().SubImage(image.Rect(sx, sy, sx+char.GetFrameProperties().FrameWidth, sy+char.GetFrameProperties().FrameHeight)).(*ebiten.Image), op)

	vector.DrawFilledCircle(screen, char.GetPosition().XCoordinate+halfSquare, char.GetPosition().YCoordinate+halfSquare, halfSquare/2, char.Colour, true)

}

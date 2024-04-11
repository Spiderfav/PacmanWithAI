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
	const strokeWidth float32 = 1

	lineColour := color.RGBA{5, 8, 173, 250}

	if squareToDraw.HasPellot {
		vector.DrawFilledCircle(screen, squareToDraw.NodePosition.XCoordinate+float32(halfSquare), squareToDraw.NodePosition.YCoordinate+float32(halfSquare), halfSquare/6, color.RGBA{255, 248, 173, 250}, true)
	}

	if squareToDraw.HasSuperPellot {
		vector.DrawFilledCircle(screen, squareToDraw.NodePosition.XCoordinate+float32(halfSquare), squareToDraw.NodePosition.YCoordinate+float32(halfSquare), halfSquare/3, color.RGBA{255, 248, 173, 250}, true)
	}

	if squareToDraw.HasWalls.HasDown {
		vector.StrokeLine(screen, squareToDraw.NodePosition.XCoordinate, squareToDraw.NodePosition.YCoordinate+float32(squareSize), squareToDraw.NodePosition.XCoordinate+float32(squareSize), squareToDraw.NodePosition.YCoordinate+float32(squareSize), strokeWidth, lineColour, false)
	}

	if squareToDraw.HasWalls.HasRight {
		vector.StrokeLine(screen, squareToDraw.NodePosition.XCoordinate+float32(squareSize), squareToDraw.NodePosition.YCoordinate, squareToDraw.NodePosition.XCoordinate+float32(squareSize), squareToDraw.NodePosition.YCoordinate+float32(squareSize), strokeWidth, lineColour, false)
	}

	if squareToDraw.HasWalls.HasLeft {
		vector.StrokeLine(screen, squareToDraw.NodePosition.XCoordinate, squareToDraw.NodePosition.YCoordinate, squareToDraw.NodePosition.XCoordinate, squareToDraw.NodePosition.YCoordinate+float32(squareSize), strokeWidth, lineColour, false)
	}

	if squareToDraw.HasWalls.HasUp {
		vector.StrokeLine(screen, squareToDraw.NodePosition.XCoordinate, squareToDraw.NodePosition.YCoordinate, squareToDraw.NodePosition.XCoordinate+float32(squareSize), squareToDraw.NodePosition.YCoordinate, strokeWidth, lineColour, false)
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

	vector.DrawFilledCircle(screen, char.GetPosition().XCoordinate+halfSquare, char.GetPosition().YCoordinate+halfSquare, halfSquare/2, char.Colour, true)

}

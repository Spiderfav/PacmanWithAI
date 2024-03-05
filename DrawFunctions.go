package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/input"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
	"golang.org/x/image/font"
)

// This function draws a given square to the screen
// It checks if the current node has a given wall, then draws it to the screen
func drawSquare(screen *ebiten.Image, squareToDraw mazegrid.MazeSquare) {
	var strokeWidth float32 = 1

	if squareToDraw.ContainsObject {
		vector.DrawFilledCircle(screen, squareToDraw.NodePosition.XCoordinate+10, squareToDraw.NodePosition.YCoordinate+10, 2, color.RGBA{255, 100, 0, 250}, true)
	}

	if squareToDraw.HasWalls.HasDown {
		vector.StrokeLine(screen, squareToDraw.NodePosition.XCoordinate, squareToDraw.NodePosition.YCoordinate+20, squareToDraw.NodePosition.XCoordinate+20, squareToDraw.NodePosition.YCoordinate+20, strokeWidth, color.Black, false)
	}

	if squareToDraw.HasWalls.HasRight {
		vector.StrokeLine(screen, squareToDraw.NodePosition.XCoordinate+20, squareToDraw.NodePosition.YCoordinate, squareToDraw.NodePosition.XCoordinate+20, squareToDraw.NodePosition.YCoordinate+20, strokeWidth, color.Black, false)
	}

	if squareToDraw.HasWalls.HasLeft {
		vector.StrokeLine(screen, squareToDraw.NodePosition.XCoordinate, squareToDraw.NodePosition.YCoordinate, squareToDraw.NodePosition.XCoordinate, squareToDraw.NodePosition.YCoordinate+20, strokeWidth, color.Black, false)
	}

	if squareToDraw.HasWalls.HasUp {
		vector.StrokeLine(screen, squareToDraw.NodePosition.XCoordinate, squareToDraw.NodePosition.YCoordinate, squareToDraw.NodePosition.XCoordinate+20, squareToDraw.NodePosition.YCoordinate, strokeWidth, color.Black, false)
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
	prevX := pathTaken[0].NodePosition.XCoordinate + 10
	prevY := pathTaken[0].NodePosition.YCoordinate + 10

	for i := 1; i < len(pathTaken); i++ {
		vector.StrokeLine(screen, prevX, prevY, pathTaken[i].NodePosition.XCoordinate+10, pathTaken[i].NodePosition.YCoordinate+10, 1, color.RGBA{0, 255, 0, 250}, false)
		prevX = pathTaken[i].NodePosition.XCoordinate + 10
		prevY = pathTaken[i].NodePosition.YCoordinate + 10

	}

}

// This function draws all the buttons to the screen for the main menu
func mainMenu(screen *ebiten.Image, g *Game) {
	// Clear the screen to white
	screen.Fill(color.White)

	text.Draw(screen, "Pacman Game", g.fontFace, (screenWidth/2)-40, (screenHeight/2)-100, color.Black)

	for i := 0; i < len(g.buttonsMenu); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(g.buttonsMenu[i].X), float64(g.buttonsMenu[i].Y))

		screen.DrawImage(g.buttonsMenu[i].Image, op)

		text.Draw(screen, g.buttonsMenu[i].Message, g.fontFace, g.buttonsMenu[i].X+10, g.buttonsMenu[i].Y+20, color.Black)
	}

}

// This functions, given an array of buttons, draws the buttons to the screen
func drawMenu(screen *ebiten.Image, arr []*input.Button, font font.Face) {

	for i := 0; i < len(arr); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(arr[i].X), float64(arr[i].Y))

		screen.DrawImage(arr[i].Image, op)

		text.Draw(screen, arr[i].Message, font, arr[i].X+10, arr[i].Y+20, color.Black)
	}
}

// This function draws the game meny to the screen
func gameMenu(screen *ebiten.Image, g *Game) {
	// 	// Clear the screen to white
	screen.Fill(color.White)
	// 	// Draw the maze to the screen
	drawMaze(screen, g)
	//OldMazeSystem(screen, g)
	backButton(screen, g)
	drawMenu(screen, g.buttonsSize, g.fontFace)
	drawMenu(screen, g.buttonsAlgo, g.fontFace)

}

// This function draws the back button to the screen
func backButton(screen *ebiten.Image, g *Game) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.buttonBack.X), float64(g.buttonBack.Y))
	screen.DrawImage(g.buttonBack.Image, op)

	text.Draw(screen, g.buttonBack.Message, g.fontFace, g.buttonBack.X+10, g.buttonBack.Y+20, color.Black)
}

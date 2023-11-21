package main

import (
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Defining the normal font for use in the program
var (
	mplusNormalFont font.Face
)

// This function draws a given square to the screen
// It checks if the current node has a given wall, then draws it to the screen
func DrawSquare(screen *ebiten.Image, squareToDraw MazeSquare) {
	var strokeWidth float32 = 1

	if squareToDraw.HasDown {
		vector.StrokeLine(screen, squareToDraw.XCoordinate, squareToDraw.YCoordinate+20, squareToDraw.XCoordinate+20, squareToDraw.YCoordinate+20, strokeWidth, color.Black, false)
	}

	if squareToDraw.HasRight {
		vector.StrokeLine(screen, squareToDraw.XCoordinate+20, squareToDraw.YCoordinate, squareToDraw.XCoordinate+20, squareToDraw.YCoordinate+20, strokeWidth, color.Black, false)
	}

	if squareToDraw.HasLeft {
		vector.StrokeLine(screen, squareToDraw.XCoordinate, squareToDraw.YCoordinate, squareToDraw.XCoordinate, squareToDraw.YCoordinate+20, strokeWidth, color.Black, false)
	}

	if squareToDraw.HasUp {
		vector.StrokeLine(screen, squareToDraw.XCoordinate, squareToDraw.YCoordinate, squareToDraw.XCoordinate+20, squareToDraw.YCoordinate, strokeWidth, color.Black, false)
	}

}

// The DrawMaze function takes the screen argument given as the screen to draw to maze to
// It draws the maze from the GameGridDFS
func DrawMaze(screen *ebiten.Image, size int) {

	// For each row and column, it looks at the walls of the block and draws the ones it has
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			DrawSquare(screen, gameGridDFS[i][j])
		}
	}
}

// This function draws lines for each path taken
func drawPathsLines(screen *ebiten.Image, pathTaken []MazeSquare) {
	prevX := pathTaken[0].XCoordinate + 10
	prevY := pathTaken[0].YCoordinate + 10

	for i := 1; i < len(pathTaken); i++ {
		vector.StrokeLine(screen, prevX, prevY, pathTaken[i].XCoordinate+10, pathTaken[i].YCoordinate+10, 1, color.RGBA{0, 255, 0, 250}, false)
		prevX = pathTaken[i].XCoordinate + 10
		prevY = pathTaken[i].YCoordinate + 10

	}

}

func drawMultiplePaths(screen *ebiten.Image, pathsTaken [][]MazeSquare) {
	for count := 0; count < len(pathsTaken); count++ {
		drawPathsLines(screen, pathsTaken[count])
	}
}

// This function draws circles with their position in the path
// It also draws the start node and end node and the total cost
func drawPaths(screen *ebiten.Image, pathTaken []MazeSquare, algo string, weight int) {

	// Here we are defining the font to be used from the general golang fonts
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    8,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})

	// For every node searched by the algorithms, draw a circle with their postion
	for i := 0; i < len(pathTaken); i++ {
		vector.DrawFilledCircle(screen, pathTaken[i].XCoordinate+10, pathTaken[i].YCoordinate+10, 2, color.RGBA{255, 0, 0, 250}, true)
		text.Draw(screen, strconv.Itoa(i), mplusNormalFont, int(pathTaken[i].XCoordinate)+10, int(pathTaken[i].YCoordinate)+10, color.RGBA{255, 0, 255, 250})

	}

	text.Draw(screen, "Path cost to desired node is "+strconv.Itoa(int(pathTaken[len(pathTaken)-1].Weight)), mplusNormalFont, 10, 10, color.RGBA{0, 0, 0, 250})
	text.Draw(screen, "Start node is "+strconv.Itoa(int(pathTaken[0].XCoordinate))+","+strconv.Itoa(int(pathTaken[0].YCoordinate)), mplusNormalFont, 10, int(gameGridDFS[len(gameGridDFS)-1][len(gameGridDFS)-1].YCoordinate)+40, color.RGBA{0, 0, 0, 250})
	text.Draw(screen, "End node is "+strconv.Itoa(int(pathTaken[len(pathTaken)-1].XCoordinate))+","+strconv.Itoa(int(pathTaken[len(pathTaken)-1].YCoordinate)), mplusNormalFont, 10, int(gameGridDFS[len(gameGridDFS)-1][len(gameGridDFS)-1].YCoordinate)+50, color.RGBA{0, 0, 0, 250})
	text.Draw(screen, "Algorithm Used: "+algo, mplusNormalFont, 10, int(gameGridDFS[len(gameGridDFS)-1][len(gameGridDFS)-1].YCoordinate)+60, color.RGBA{0, 0, 0, 250})
	text.Draw(screen, "Total Weight: "+strconv.Itoa(weight), mplusNormalFont, 10, int(gameGridDFS[len(gameGridDFS)-1][len(gameGridDFS)-1].YCoordinate)+70, color.RGBA{0, 0, 0, 250})

}

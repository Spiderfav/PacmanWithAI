package main

import (
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Defining the normal font for use in the program
var (
	mplusNormalFont font.Face
)

// This function draws a given square to the screen
// It checks if the current node has a given wall, then draws it to the screen
func drawSquare(screen *ebiten.Image, squareToDraw mazegrid.MazeSquare) {
	var strokeWidth float32 = 1

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
func drawMaze(screen *ebiten.Image, size int) {

	// For each row and column, it looks at the walls of the block and draws the ones it has
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			drawSquare(screen, gameGridDFS[i][j])
		}
	}
}

// This function draws lines for each path taken
func drawPathsLines(screen *ebiten.Image, pathTaken []mazegrid.MazeSquare) {
	prevX := pathTaken[0].NodePosition.XCoordinate + 10
	prevY := pathTaken[0].NodePosition.YCoordinate + 10

	for i := 1; i < len(pathTaken); i++ {
		vector.StrokeLine(screen, prevX, prevY, pathTaken[i].NodePosition.XCoordinate+10, pathTaken[i].NodePosition.YCoordinate+10, 1, color.RGBA{0, 255, 0, 250}, false)
		prevX = pathTaken[i].NodePosition.XCoordinate + 10
		prevY = pathTaken[i].NodePosition.YCoordinate + 10

	}

}

func drawMultiplePaths(screen *ebiten.Image, pathsTaken [][]mazegrid.MazeSquare) {
	for count := 0; count < len(pathsTaken); count++ {
		drawPathsLines(screen, pathsTaken[count])
	}
}

// This function draws circles with their position in the path
// It also draws the start node and end node and the total cost
func drawPaths(screen *ebiten.Image, pathTaken []mazegrid.MazeSquare, algo string, weight int) {

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
		vector.DrawFilledCircle(screen, pathTaken[i].NodePosition.XCoordinate+10, pathTaken[i].NodePosition.YCoordinate+10, 2, color.RGBA{255, 0, 0, 250}, true)
		text.Draw(screen, strconv.Itoa(i), mplusNormalFont, int(pathTaken[i].NodePosition.XCoordinate)+10, int(pathTaken[i].NodePosition.YCoordinate)+10, color.RGBA{255, 0, 255, 250})

	}

	text.Draw(screen, "Path cost to desired node is "+strconv.Itoa(int(pathTaken[len(pathTaken)-1].Weight)), mplusNormalFont, 10, 10, color.RGBA{0, 0, 0, 250})
	text.Draw(screen, "Start node is "+strconv.Itoa(int(pathTaken[0].NodePosition.XCoordinate))+","+strconv.Itoa(int(pathTaken[0].NodePosition.YCoordinate)), mplusNormalFont, 10, int(gameGridDFS[len(gameGridDFS)-1][len(gameGridDFS)-1].NodePosition.YCoordinate)+40, color.RGBA{0, 0, 0, 250})
	text.Draw(screen, "End node is "+strconv.Itoa(int(pathTaken[len(pathTaken)-1].NodePosition.XCoordinate))+","+strconv.Itoa(int(pathTaken[len(pathTaken)-1].NodePosition.YCoordinate)), mplusNormalFont, 10, int(gameGridDFS[len(gameGridDFS)-1][len(gameGridDFS)-1].NodePosition.YCoordinate)+50, color.RGBA{0, 0, 0, 250})
	text.Draw(screen, "Algorithm Used: "+algo, mplusNormalFont, 10, int(gameGridDFS[len(gameGridDFS)-1][len(gameGridDFS)-1].NodePosition.YCoordinate)+60, color.RGBA{0, 0, 0, 250})
	text.Draw(screen, "Total Weight: "+strconv.Itoa(weight), mplusNormalFont, 10, int(gameGridDFS[len(gameGridDFS)-1][len(gameGridDFS)-1].NodePosition.YCoordinate)+70, color.RGBA{0, 0, 0, 250})

}
func mainMenu(screen *ebiten.Image, g *Game) {
	// Clear the screen to white
	screen.Fill(color.White)

	text.Draw(screen, "Pacman Game", g.fontFace, (screenWidth/2)-40, (screenHeight/2)-100, color.Black)

	for i := 0; i < 3; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(g.buttonsMenu[i].x), float64(g.buttonsMenu[i].y))

		screen.DrawImage(g.buttonsMenu[i].image, op)

		text.Draw(screen, g.buttonsMenu[i].message, g.fontFace, g.buttonsMenu[i].x+10, g.buttonsMenu[i].y+20, color.Black)
	}

}

func sizeMenu(screen *ebiten.Image, g *Game) {

	for i := 0; i < len(g.buttonsSize); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(g.buttonsSize[i].x), float64(g.buttonsSize[i].y))

		screen.DrawImage(g.buttonsSize[i].image, op)

		text.Draw(screen, g.buttonsSize[i].message, g.fontFace, g.buttonsSize[i].x+10, g.buttonsSize[i].y+20, color.Black)
	}
}

func algoMenu(screen *ebiten.Image, g *Game) {

	for i := 0; i < len(g.buttonsAlgo); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(g.buttonsAlgo[i].x), float64(g.buttonsAlgo[i].y))

		screen.DrawImage(g.buttonsAlgo[i].image, op)

		text.Draw(screen, g.buttonsAlgo[i].message, g.fontFace, g.buttonsAlgo[i].x+10, g.buttonsAlgo[i].y+20, color.Black)
	}
}

func backButton(screen *ebiten.Image, g *Game) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.buttonBack.x), float64(g.buttonBack.y))
	screen.DrawImage(g.buttonBack.image, op)

	text.Draw(screen, g.buttonBack.message, g.fontFace, g.buttonBack.x+10, g.buttonBack.y+20, color.Black)
}

func OldMazeSystem(screen *ebiten.Image, g *Game) {
	// Clear the screen to white
	screen.Fill(color.White)
	// Draw the maze to the screen
	drawMaze(screen, mazeSize)

	if whichPath == 0 {
		// Clear the screen to white
		screen.Fill(color.White)

		// Draw the maze to the screen
		drawMaze(screen, mazeSize)

		// Draw Dijkstra's Path to the screen
		drawPaths(screen, dijkstrasPath, "Dijstra", weightDijkstras)
		drawPathsLines(screen, absolutePathDijkstras)

	} else if whichPath == 1 {
		// Clear the screen to white
		screen.Fill(color.White)

		// Draw the maze to the screen
		drawMaze(screen, mazeSize)

		// Draw A*'s Path to the screen
		drawPaths(screen, aStarPath, "A Star", weigthAStar)
		drawPathsLines(screen, absolutePathAStar)

	} else if whichPath == 2 {
		// Clear the screen to white
		screen.Fill(color.White)

		// Draw the maze to the screen
		drawMaze(screen, mazeSize)
		drawPaths(screen, graph, "Graph Method", 10)
		drawMultiplePaths(screen, graphPaths)

	} else if whichPath == 4 {
		// Clear the screen to white
		screen.Fill(color.White)

		// Draw the maze to the screen
		drawMaze(screen, mazeSize)

		// Draw Solution Path to the screen
		drawPathsLines(screen, absolutePathAStar)

	}
}

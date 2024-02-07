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

// func drawMultiplePaths(screen *ebiten.Image, pathsTaken [][]mazegrid.MazeSquare) {
// 	for count := 0; count < len(pathsTaken); count++ {
// 		drawPathsLines(screen, pathsTaken[count])
// 	}
// }

// This function draws circles with their position in the path
// It also draws the start node and end node and the total cost
// func drawPaths(screen *ebiten.Image, pathTaken []mazegrid.MazeSquare, algo string, weight int) {

// 	// For every node searched by the algorithms, draw a circle with their postion
// 	for i := 0; i < len(pathTaken); i++ {
// 		vector.DrawFilledCircle(screen, pathTaken[i].NodePosition.XCoordinate+10, pathTaken[i].NodePosition.YCoordinate+10, 2, color.RGBA{255, 0, 0, 250}, true)
// 		text.Draw(screen, strconv.Itoa(i), basicfont.Face7x13, int(pathTaken[i].NodePosition.XCoordinate)+10, int(pathTaken[i].NodePosition.YCoordinate)+10, color.RGBA{255, 0, 255, 250})

// 	}

// 	text.Draw(screen, "Path cost to desired node is "+strconv.Itoa(int(pathTaken[len(pathTaken)-1].Weight)), basicfont.Face7x13, 10, 10, color.RGBA{0, 0, 0, 250})
// 	text.Draw(screen, "Start node is "+strconv.Itoa(int(pathTaken[0].NodePosition.XCoordinate))+","+strconv.Itoa(int(pathTaken[0].NodePosition.YCoordinate)), basicfont.Face7x13, 10, int(gameGridDFS[len(gameGridDFS)-1][len(gameGridDFS)-1].NodePosition.YCoordinate)+40, color.RGBA{0, 0, 0, 250})
// 	text.Draw(screen, "End node is "+strconv.Itoa(int(pathTaken[len(pathTaken)-1].NodePosition.XCoordinate))+","+strconv.Itoa(int(pathTaken[len(pathTaken)-1].NodePosition.YCoordinate)), basicfont.Face7x13, 10, int(gameGridDFS[len(gameGridDFS)-1][len(gameGridDFS)-1].NodePosition.YCoordinate)+50, color.RGBA{0, 0, 0, 250})
// 	text.Draw(screen, "Algorithm Used: "+algo, basicfont.Face7x13, 10, int(gameGridDFS[len(gameGridDFS)-1][len(gameGridDFS)-1].NodePosition.YCoordinate)+60, color.RGBA{0, 0, 0, 250})
// 	text.Draw(screen, "Total Weight: "+strconv.Itoa(weight), basicfont.Face7x13, 10, int(gameGridDFS[len(gameGridDFS)-1][len(gameGridDFS)-1].NodePosition.YCoordinate)+70, color.RGBA{0, 0, 0, 250})

// }

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

func drawMenu(screen *ebiten.Image, arr []*input.Button, font font.Face) {

	for i := 0; i < len(arr); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(arr[i].X), float64(arr[i].Y))

		screen.DrawImage(arr[i].Image, op)

		text.Draw(screen, arr[i].Message, font, arr[i].X+10, arr[i].Y+20, color.Black)
	}
}

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

func backButton(screen *ebiten.Image, g *Game) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.buttonBack.X), float64(g.buttonBack.Y))
	screen.DrawImage(g.buttonBack.Image, op)

	text.Draw(screen, g.buttonBack.Message, g.fontFace, g.buttonBack.X+10, g.buttonBack.Y+20, color.Black)
}

// func OldMazeSystem(screen *ebiten.Image, g *Game) {
// 	// Clear the screen to white
// 	screen.Fill(color.White)
// 	// Draw the maze to the screen
// 	drawMaze(screen, mazeSize)

// 	if whichPath == 0 {

// 		// Draw Dijkstra's Path to the screen
// 		drawPaths(screen, dijkstrasPath, "Dijstra", weightDijkstras)
// 		drawPathsLines(screen, absolutePathDijkstras)
// 		return

// 	} else if whichPath == 1 {

// 		// Draw A*'s Path to the screen
// 		drawPaths(screen, aStarPath, "A Star", weigthAStar)
// 		drawPathsLines(screen, absolutePathAStar)
// 		return

// 	} else if whichPath == 2 {

// 		drawPaths(screen, graph, "Graph Method", 10)
// 		drawMultiplePaths(screen, graphPaths)
// 		return

// 	} else if whichPath == 4 {

// 		// Draw Solution Path to the screen
// 		drawPathsLines(screen, absolutePathAStar)
// 		return

// 	}
// }

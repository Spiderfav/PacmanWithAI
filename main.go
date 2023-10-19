package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type MazeSquare struct {
	XCoordinate float32
	YCoordinate float32
	Left        *MazeSquare
	HasLeft     bool
	Down        *MazeSquare
	HasDown     bool
	Right       *MazeSquare
	HasRight    bool
	Up          *MazeSquare
	HasUp       bool
	Visited     bool
}

type Game struct{}

var gameGridDFS [8][8]MazeSquare = DFS()

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.White)
	// DFS(8, gameGrid)

	DrawMaze(screen)

}

func DrawMaze(screen *ebiten.Image) {

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			DrawSquare(screen, gameGridDFS[i][j])
		}
	}
}

func CreateMaze() [8][8]MazeSquare {
	var gameGrid [8][8]MazeSquare

	var x, y float32

	var squareLengthX, squareLengthY float32 = 20, 20

	for y = 0; y < 8; y++ {

		for x = 0; x < 8; x++ {

			// Using i + 1 and j + 1 as this is calculating the square size and as it starts by 0, we need to add one to the normal counter
			var square = MazeSquare{squareLengthX * (x + 1), squareLengthY * (y + 1), nil, true, nil, true, nil, true, nil, true, false}
			//square.DrawSquare(screen, squareLengthX*(i+1), squareLengthY*(j+1))

			gameGrid[int(y)][int(x)] = square

			if x > 0 {
				//fmt.Println("Square ", square)
				gameGrid[int(y)][int(x)].Left = &gameGrid[int(y)][int(x-1)]
				//fmt.Println("Square next ", square)
				gameGrid[int(y)][int(x-1)].Right = &gameGrid[int(y)][int(x)]
			}

			if y > 0 {
				gameGrid[int(y)][int(x)].Up = &gameGrid[int(y-1)][int(x)]
				gameGrid[int(y-1)][int(x)].Down = &gameGrid[int(y)][int(x)]

			}

		}

	}

	return gameGrid
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Single Agent Maze!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

/*

The & Operator

	& goes in front of a variable when you want to get that variable's memory address.

The * Operator

	* goes in front of a variable that holds a memory address and resolves it (it is therefore the counterpart to the & operator).
	It goes and gets the thing that the pointer was pointing at

*/

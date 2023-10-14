package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type MazeSquare struct {
	XCoordinate float32
	YCoordinate float32
	Left        *MazeSquare
	Down        *MazeSquare
	Right       *MazeSquare
	Up          *MazeSquare
}

var gameGrid [8][8]MazeSquare

func (square MazeSquare) DrawSquare(screen *ebiten.Image, x float32, y float32) {

	var strokeWidth float32 = 1

	vector.StrokeLine(screen, x, y, x+20, y, strokeWidth, color.Black, false)
	vector.StrokeLine(screen, x+20, y, x+20, y+20, strokeWidth, color.Black, false)

	vector.StrokeLine(screen, x+20, y+20, x, y+20, strokeWidth, color.Black, false)
	vector.StrokeLine(screen, x, y+20, x, y, strokeWidth, color.Black, false)

}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.White)
	DrawMaze(screen)
	breakLine(gameGrid[3][3], screen)
}

func DrawMaze(screen *ebiten.Image) {

	var i, j float32

	var squareLengthX, squareLengthY float32 = 20, 20

	for j = 0; j < 8; j += 1 {

		for i = 0; i < 8; i += 1 {

			// Using i + 1 and j + 1 as this is calculating the square size and as it starts by 0, we need to add one to the normal counter
			var square = MazeSquare{squareLengthX * (i + 1), squareLengthY * (j + 1), nil, nil, nil, nil}
			square.DrawSquare(screen, squareLengthX*(i+1), squareLengthY*(j+1))

			if i > 0 {
				square.Left = &gameGrid[int(j)][int(i-1)]
				gameGrid[int(j)][int(i-1)].Right = &square
			}

			if j > 0 || (i == 0 && j > 0) {
				square.Up = &gameGrid[int(j-1)][int(i)]
				gameGrid[int(j-1)][int(i)].Down = &square

			}

			gameGrid[int(i)][int(j)] = square

		}
	}

	//fmt.Println(gameGrid[3][3])
	// fmt.Println(gameGrid[0][0].Down.YCoordinate)
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

	* goes in front of a variable that holds a memory address and resolves it (it is therefore the counterpart to the & operator). It goes and gets the thing that the pointer was pointing at

*/

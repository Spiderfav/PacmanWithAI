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
}

func DrawMaze(screen *ebiten.Image) {

	var i, j float32

	var prevSquare MazeSquare

	var squareLengthX, squareLengthY float32 = 20, 20

	for j = 1; j < 9; j += 1 {

		for i = 1; i < 9; i += 1 {

			var square = MazeSquare{squareLengthX * i, squareLengthY * j, nil, nil, nil, nil}
			square.DrawSquare(screen, squareLengthX*i, squareLengthY*j)

			if (prevSquare != MazeSquare{}) {
				square.Left = &prevSquare
				prevSquare.Right = &square
			}

			prevSquare = square

			gameGrid[int(j)-1][int(i-1)] = square

		}
	}

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

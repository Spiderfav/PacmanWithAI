package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type MazePoints struct {
	Start MazeSquare
	End   MazeSquare
}
type MazeSquare struct {
	//X     float64
	Left  *MazeSquare
	Down  *MazeSquare
	Right *MazeSquare
	Up    *MazeSquare
}

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

	var i float32

	var maze MazePoints

	var prevSquare MazeSquare

	var squareLength float32 = 20

	for i = 1; i < 9; i += 1 {
		fmt.Println(i)

		var square = MazeSquare{nil, nil, nil, nil}
		square.DrawSquare(screen, squareLength*i, squareLength)

		if (prevSquare != MazeSquare{}) {
			square.Left = &prevSquare
			prevSquare.Right = &square
		}

		prevSquare = square

		if i == 1 {
			maze.Start = square
		}

		if i == 8 {
			maze.End = square
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

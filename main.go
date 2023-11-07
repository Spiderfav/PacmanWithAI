package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var gameGridDFS [8][8]MazeSquare = DFS()

var shortestPath1 = dijkstras(&gameGridDFS, 20, 20, 160, 160)

var shortestPath2 = aStar(&gameGridDFS, 20, 20, 160, 160)

func (g *Game) Update() error {
	return nil
}

// This function is called every second to update what is drawn on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen to white
	screen.Fill(color.White)

	// Draw the maze to the screen
	DrawMaze(screen)
	drawPaths(screen, shortestPath2)
	drawPathsLines(screen, shortestPath1)

}

// The DrawMaze function takes the screen argument given as the screen to draw to maze to
// It draws the maze from the GameGridDFS
func DrawMaze(screen *ebiten.Image) {

	// For each row and column, it looks at the walls of the block and draws the ones it has
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			DrawSquare(screen, gameGridDFS[i][j])
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {

	fmt.Println("Size of Dijkstras:", len(shortestPath1))
	fmt.Println("Size of A*:", len(shortestPath2))

	ebiten.SetWindowSize(1280, 720)
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

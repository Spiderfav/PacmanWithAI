package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

var mazeSize = 38

var gameGridDFS [][]MazeSquare = DFS(mazeSize)

var shortestPath1 = dijkstras(gameGridDFS, 20, 20, 20*mazeSize, 20*mazeSize)

var shortestPath2 = aStar(gameGridDFS, 20, 20, 20*mazeSize, 20*mazeSize)

var absolutePath1 = absolutePath(shortestPath1)

var absolutePath2 = absolutePath(shortestPath2)
var whichPath = 0

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		whichPath = 0

	} else if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		whichPath = 1
	}

	return nil
}

// This function is called every second to update what is drawn on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen to white
	screen.Fill(color.White)
	// Draw the maze to the screen
	DrawMaze(screen, mazeSize)

	if whichPath == 0 {
		// Clear the screen to white
		screen.Fill(color.White)

		// Draw the maze to the screen
		DrawMaze(screen, mazeSize)

		// Draw Dijkstra's Path to the screen
		drawPaths(screen, shortestPath1, "Dijstra")
		drawPathsLines(screen, absolutePath1)

	} else if whichPath == 1 {
		// Clear the screen to white
		screen.Fill(color.White)

		// Draw the maze to the screen
		DrawMaze(screen, mazeSize)

		// Draw A*'s Path to the screen
		drawPaths(screen, shortestPath2, "A Star")
		drawPathsLines(screen, absolutePath2)

	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func main() {

	fmt.Println("Size of Dijkstras:", len(shortestPath1))
	fmt.Println("Size of A*:", len(shortestPath2))

	fmt.Println("Size of absolute path", len(absolutePath1))

	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Single Agent Maze!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

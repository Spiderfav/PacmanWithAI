package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

var mazeSizeOriginal = 8

var gameGridDFS [][]MazeSquare = DFS(mazeSizeOriginal)

var mazeSize = len(gameGridDFS[0])

var dijkstrasPath = dijkstras(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)

var aStarPath = aStar(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)

var absolutePathDijkstras = absolutePath(dijkstrasPath)

var absolutePathAStar = absolutePath(aStarPath)

var whichPath = 3

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		changeMazeSize(mazeSizeOriginal)
	}

	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		changeMazeSize(mazeSizeOriginal * 2)

	}

	if inpututil.IsKeyJustPressed(ebiten.Key3) {
		changeMazeSize((mazeSizeOriginal * 2) * 2)

	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		whichPath = 0

	} else if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		whichPath = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		whichPath = 3
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
		drawPaths(screen, dijkstrasPath, "Dijstra")
		drawPathsLines(screen, absolutePathDijkstras)

	} else if whichPath == 1 {
		// Clear the screen to white
		screen.Fill(color.White)

		// Draw the maze to the screen
		DrawMaze(screen, mazeSize)

		// Draw A*'s Path to the screen
		drawPaths(screen, aStarPath, "A Star")
		drawPathsLines(screen, absolutePathAStar)

	} else if whichPath == 4 {
		// Clear the screen to white
		screen.Fill(color.White)

		// Draw the maze to the screen
		DrawMaze(screen, mazeSize)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func main() {

	fmt.Println("Size of Dijkstras:", len(dijkstrasPath))
	fmt.Println("Size of A*:", len(aStarPath))

	fmt.Println("Size of absolute path", len(absolutePathDijkstras))

	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Single Agent Maze!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func changeMazeSize(newSize int) {
	gameGridDFS = DFS(newSize)
	dijkstrasPath = dijkstras(gameGridDFS, 20, 20, 20*newSize, 20*newSize)
	aStarPath = aStar(gameGridDFS, 20, 20, 20*newSize, 20*newSize)
	absolutePathDijkstras = absolutePath(dijkstrasPath)
	absolutePathAStar = absolutePath(aStarPath)
	mazeSize = newSize
	whichPath = 4
}

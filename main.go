package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/generation"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

type Game struct{}

var mazeSizeOriginal = 8
var gameGridDFS [][]mazegrid.MazeSquare = generation.DFS(mazeSizeOriginal)
var mazeSize = len(gameGridDFS[0])

var dijkstrasPath = algorithms.Dijkstras(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
var absolutePathDijkstras, weightDijkstras = algorithms.AbsolutePath(dijkstrasPath)

var aStarPath = algorithms.AStar(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
var absolutePathAStar, weigthAStar = algorithms.AbsolutePath(aStarPath)

var graph = generation.MazeToGraph(gameGridDFS, 20, 20, float32(20*mazeSizeOriginal), float32(20*mazeSizeOriginal))
var graphPaths = generation.AllPaths(gameGridDFS, graph)

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
		whichPath = 2

	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		whichPath = 3
	}

	return nil
}

// This function is called every second to update what is drawn on the screen
func (g *Game) Draw(screen *ebiten.Image) {
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

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func main() {

	fmt.Println("Size of Dijkstras:", len(dijkstrasPath))
	fmt.Println("Size of A*:", len(aStarPath))

	fmt.Println("Size of absolute path", len(absolutePathDijkstras))
	fmt.Println(" ")

	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Single Agent Maze!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func changeMazeSize(newSize int) {
	gameGridDFS = generation.DFS(newSize)
	dijkstrasPath = algorithms.Dijkstras(gameGridDFS, 20, 20, 20*newSize, 20*newSize)
	aStarPath = algorithms.AStar(gameGridDFS, 20, 20, 20*newSize, 20*newSize)
	absolutePathDijkstras, weightDijkstras = algorithms.AbsolutePath(dijkstrasPath)
	absolutePathAStar, weigthAStar = algorithms.AbsolutePath(aStarPath)
	mazeSize = newSize
	whichPath = 3
	graph = generation.MazeToGraph(gameGridDFS, 20, 20, float32(20*newSize), float32(20*newSize))
	graphPaths = generation.AllPaths(gameGridDFS, graph)
}

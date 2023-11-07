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

package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var strokeWidth float32 = 1

func DFS(gridSize int, gameGrid [][]MazeSquare) {

	var visitedGrid [][]int

	startPoint := rand.Intn(gridSize)
	startNode := gameGrid[startPoint][startPoint]
	visitedGrid[startPoint][startPoint] = 1

	//var stack []MazeSquare

	moveOptions := [4]MazeSquare{*startNode.Left, *startNode.Down, *startNode.Right, *startNode.Up}

	//stack = append(stack, startNode)

	direction := moveOptions[rand.Intn(3)]

	for (direction == MazeSquare{}) {
		direction = moveOptions[rand.Intn(3)]
	}

}

func breakLine(current MazeSquare, screen *ebiten.Image) {

	// fmt.Println("Current: ", current)
	// fmt.Println("Current X :", current.XCoordinate)
	// fmt.Println("Current Y :", current.YCoordinate)
	// fmt.Println("Left Y :", current.Left)

	// fmt.Println("Map Works? ", mazeMap[3][4])
	// fmt.Println("Map Saves? ", mazeMap[3][4].Left)

	//vector.StrokeLine(screen, 80, 80, 100, 80, strokeWidth, color.White, false)

	vector.StrokeLine(screen, current.XCoordinate, current.YCoordinate, current.Left.XCoordinate, current.Left.YCoordinate, strokeWidth, color.White, false)
}

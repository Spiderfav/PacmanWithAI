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
	var stack []MazeSquare

	// Randomly selecting a node
	startPoint := rand.Intn(gridSize)
	startNode := gameGrid[startPoint][startPoint]

	// Appending node to stack
	stack = append(stack, startNode)

	// Marking node as visited
	visitedGrid[startPoint][startPoint] = 1

	// Choose random direction to go in
	moveOptions := [4]MazeSquare{*startNode.Left, *startNode.Down, *startNode.Right, *startNode.Up}

	direction := moveOptions[rand.Intn(3)]

	// If direction picked does not have a node, pick another direction
	for (direction == MazeSquare{}) {
		direction = moveOptions[rand.Intn(3)]
	}

}

func breakLine(current MazeSquare, screen *ebiten.Image, directionToBreak int) {

	// 0 = Left, 1 = Down, 2 = Right, 3 = Up

	switch directionToBreak {
	case 0:
		vector.StrokeLine(screen, current.XCoordinate, current.YCoordinate, current.Down.XCoordinate, current.Down.YCoordinate, strokeWidth, color.White, false)

	case 1:
		vector.StrokeLine(screen, current.Down.XCoordinate, current.Down.YCoordinate, current.Down.Right.XCoordinate, current.Down.Right.YCoordinate, strokeWidth, color.White, false)

	case 2:
		vector.StrokeLine(screen, current.Right.XCoordinate, current.Right.YCoordinate, current.Right.Down.XCoordinate, current.Right.Down.YCoordinate, strokeWidth, color.White, false)

	case 3:
		vector.StrokeLine(screen, current.XCoordinate, current.YCoordinate, current.Right.XCoordinate, current.Right.YCoordinate, strokeWidth, color.White, false)

	}

}

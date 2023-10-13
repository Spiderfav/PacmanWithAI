package main

import "math/rand"

func DFS(gridSize int, gameGrid [][]MazeSquare) {

	var visitedGrid [][]int

	startPoint := rand.Intn(gridSize)
	startNode := gameGrid[startPoint][startPoint]
	visitedGrid[startPoint][startPoint] = 1

	var stack []MazeSquare

	stack = append(stack, startNode)

}

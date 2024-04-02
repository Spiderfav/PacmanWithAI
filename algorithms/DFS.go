package algorithms

import (
	"fmt"
	"time"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/file"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// DFS aims to find a path to any given point by traversing through all nodes in the graph until the node is reached
func DFSearch(gameGridDFS [][]mazegrid.MazeSquare, startX int, startY int, finishX int, finishY int, squareSize int) []mazegrid.MazeSquare {
	// Marking every node unvisited
	MarkUnvisited(gameGridDFS, false)

	start := time.Now()
	fmt.Println("Mem usage before:")
	file.PrintMemUsage()

	// Initialize the stack
	var stack []mazegrid.MazeSquare

	endNode := &gameGridDFS[(finishY/squareSize)-1][(finishX/squareSize)-1]

	// Start from the starting node
	startNode := &gameGridDFS[(startY/squareSize)-1][(startX/squareSize)-1]
	startNode.Visited = true
	stack = append(stack, *startNode)

	var pathTaken []mazegrid.MazeSquare // Stores the path found

	for len(stack) > 0 {
		// Pop a node from the stack
		currentNodeIndex := len(stack) - 1
		currentNode := stack[currentNodeIndex]
		stack = stack[:currentNodeIndex]

		pathTaken = append(pathTaken, currentNode)

		// Check if this is the finish node
		if currentNode == *endNode {
			break
		}

		// Get all the possible moves from that given square
		possibleMoves := getPossibleMoves(gameGridDFS, currentNode.NodePosition, squareSize)

		// From those given moves, check which ones have already been visited and add them to the stack
		for i := 0; i < len(possibleMoves); i++ {
			nodeToTest := &gameGridDFS[(int(possibleMoves[i].YCoordinate)/squareSize)-1][(int(possibleMoves[i].XCoordinate)/squareSize - 1)]

			if !nodeToTest.Visited {

				nodeToTest.Visited = true

				stack = append(stack, *nodeToTest)

			}
		}

		// Mark the current node as visited
		currentNode.Visited = true
	}

	elapsed := time.Since(start)
	fmt.Printf("DFS took %s", elapsed)
	fmt.Println("\nDFS Concluded")
	fmt.Println(" ")

	fmt.Println("Mem usage after:")
	file.PrintMemUsage()
	fmt.Println(" ")

	return pathTaken
}

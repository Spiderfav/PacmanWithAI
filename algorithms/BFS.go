package algorithms

import (
	"fmt"
	"time"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/file"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// BFS aims to find the shortest path to any given point by traversing through all nodes in the graph until the node is reached
// The maze must be built with type mazegrid.Mazesquare
func BFS(gameGridDFS [][]mazegrid.MazeSquare, startX int, startY int, finishX int, finishY int, squareSize int) []mazegrid.MazeSquare {
	// Marking every node unvisited
	MarkUnvisited(gameGridDFS, false)

	start := time.Now() // This is used to time how long the function took to execute
	fmt.Println("Mem usage before:")
	file.PrintMemUsage()

	// Initialize the FIFO queue
	var queue []*mazegrid.MazeSquare

	endNode := &gameGridDFS[(finishY/squareSize)-1][(finishX/squareSize)-1]

	// Start from the starting node
	startNode := &gameGridDFS[(startY/squareSize)-1][(startX/squareSize)-1]
	startNode.Visited = true

	queue = append(queue, startNode)

	// Store each node's predecessor for path reconstruction
	predecessor := make(map[*mazegrid.MazeSquare]*mazegrid.MazeSquare)

	for len(queue) > 0 {

		// Dequeue a node from the front of the queue
		currentNode := queue[0]
		queue = queue[1:]

		// Check if this is the end node
		if currentNode == endNode {
			break
		}

		// Get all the possible moves from that given square
		possibleMoves := getPossibleMoves(gameGridDFS, currentNode.NodePosition, squareSize)

		// From those given moves, check which ones have already been visited and add them to the FIFO queue
		for i := 0; i < len(possibleMoves); i++ {
			nodeToTest := &gameGridDFS[(int(possibleMoves[i].YCoordinate)/squareSize)-1][(int(possibleMoves[i].XCoordinate)/squareSize - 1)]

			if !nodeToTest.Visited {

				nodeToTest.Visited = true

				queue = append(queue, nodeToTest)

				predecessor[nodeToTest] = currentNode
			}
		}

	}

	// Reconstruct path
	pathTaken := PathReconstructor(startNode, endNode, predecessor)

	elapsed := time.Since(start)
	fmt.Printf("BFS took %s", elapsed)
	fmt.Println("\nBFS Concluded")
	fmt.Println(" ")

	fmt.Println("Mem usage after:")
	file.PrintMemUsage()
	fmt.Println(" ")

	return pathTaken
}

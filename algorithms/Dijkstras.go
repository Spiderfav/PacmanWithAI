package algorithms

import (
	"container/heap"
	"fmt"
	"time"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/file"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// Dijkstras uses Dijkstras Algorithm to find the shortest path from one node to another in a given maze
// The maze must be built with type mazegrid.Mazesquare
func Dijkstras(gameGrid [][]mazegrid.MazeSquare, startX int, startY int, finishX int, finishY int, squareSize int) []mazegrid.MazeSquare {
	// Marking every node unvisited
	MarkUnvisited(gameGrid, true)

	start := time.Now() // This is used to time how long the function took to execute
	fmt.Println("Mem usage before:")
	file.PrintMemUsage()

	priorityQueue := make(PriorityQueue, 0)
	heap.Init(&priorityQueue)

	// Store each node's predecessor for path reconstruction
	predecessor := make(map[*mazegrid.MazeSquare]*mazegrid.MazeSquare)

	startNode := &gameGrid[(startY/squareSize)-1][(startX/squareSize)-1]
	startNode.Weight = 0.0

	endNode := &gameGrid[(finishY/squareSize)-1][(finishX/squareSize)-1]

	heap.Push(&priorityQueue, &PriorityNode{node: startNode, priority: startNode.Weight})

	for len(priorityQueue) > 0 {

		currentNode := heap.Pop(&priorityQueue).(*PriorityNode).node //Asserting the type from the pop

		// Check if this is the end node
		if currentNode == endNode {
			break
		}

		// Get all the possible moves from that given square
		possibleMoves := getPossibleMoves(gameGrid, currentNode.NodePosition, squareSize)

		// From those given moves, check which ones have already been visited and add them to the FIFO queue
		for i := 0; i < len(possibleMoves); i++ {
			nodeToTest := &gameGrid[(int(possibleMoves[i].YCoordinate)/squareSize)-1][(int(possibleMoves[i].XCoordinate)/squareSize - 1)]

			if !nodeToTest.Visited {

				nodeToTest.Weight = currentNode.Weight + 1
				nodeToTest.Visited = true
				heap.Push(&priorityQueue, &PriorityNode{node: nodeToTest, priority: nodeToTest.Weight})
				predecessor[nodeToTest] = currentNode

			}

		}

		currentNode.Visited = true

	}

	// Reconstruct path
	pathTaken := PathReconstructor(startNode, endNode, predecessor)

	elapsed := time.Since(start)
	fmt.Printf("Dijkstra's took %s", elapsed)
	fmt.Println("\nDijkstra Concluded")
	fmt.Println(" ")

	fmt.Println("Mem usage after:")
	file.PrintMemUsage()
	fmt.Println(" ")

	// Returns the path that the algorithm took to get from the start to the finish
	return pathTaken
}

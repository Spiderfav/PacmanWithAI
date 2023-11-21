package algorithms

import (
	"fmt"
	"time"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// This function uses Dijkstras Algorithm to find the shortest path from one node to another in a given maze
func Dijkstras(gameGridDFS [][]mazegrid.MazeSquare, startX int, startY int, finishX int, finishY int) []mazegrid.MazeSquare {
	start := time.Now()

	// Storing the original start values
	originalStartX := startX
	originalStartY := startY

	// Marking every node unvisited
	MarkUnvisited(gameGridDFS)
	AddWeights(gameGridDFS, 100)

	var pathTaken []mazegrid.MazeSquare

	prevWeight := 0
	var nodePrevWeights []int

	var splitNodes []mazegrid.MazeSquare

	// Assigning the first node a weight of 0
	gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = 0

	// While the end node of the grid has not been visited

	for !gameGridDFS[int(finishX/20)-1][int(finishY/20)-1].Visited {

		// Assigning a new weight to the current node only if it is not the starting point
		if gameGridDFS[int(startX/20)-1][int(startY/20)-1] != gameGridDFS[(originalStartX/20)-1][(originalStartY/20)-1] {
			prevWeight += 1
			gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = prevWeight + gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight
		}

		// Mark the current node as visited and add the node to the array of nodes for the path taken
		gameGridDFS[int(startX/20)-1][int(startY/20)-1].Visited = true
		pathTaken = append(pathTaken, gameGridDFS[int(startX/20)-1][int(startY/20)-1])

		// This if block checks if the current node has any neighbours and if so, adds them all sequentially to an array
		// It also stores the current weight at the given node for backtracking (that way the weight is correct)
		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasDown && !gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].Visited {
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1+1][int(startY/20)-1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)
		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasUp && !gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].Visited {
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1-1][int(startY/20)-1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasLeft && !gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].Visited {
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1][int(startY/20)-1-1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasRight && !gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].Visited {
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1][int(startY/20)-1+1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)

		}

		// If no path was possible from the current node, try a previous found neighbour of a node and set that as the new start
		if len(splitNodes) != 0 {
			nodePopped := splitNodes[len(splitNodes)-1]
			splitNodes = splitNodes[:len(splitNodes)-1]

			prevWeight = nodePrevWeights[len(nodePrevWeights)-1]
			nodePrevWeights = nodePrevWeights[:len(nodePrevWeights)-1]

			startY = int(nodePopped.XCoordinate)
			startX = int(nodePopped.YCoordinate)
		}

	}

	elapsed := time.Since(start)
	fmt.Printf("Dijkstra's took %s", elapsed)
	fmt.Println("\nDijkstra Concluded")
	fmt.Println(" ")

	// Returns the path that the algorithm took to get from the start to the finish
	return pathTaken
}

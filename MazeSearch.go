package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

// This function uses the A* Algorithm to find the shortest path from one node to another in a given maze
func aStar(gameGridDFS *[8][8]MazeSquare, startX int, startY int, finishX int, finishY int) []MazeSquare {
	start := time.Now()

	// Storing the original start values
	originalStartX := startX
	originalStartY := startY

	// Marking every node unvisited
	markUnvisited(gameGridDFS)

	var bestPath []MazeSquare

	prevWeight := 0
	var nodePrevWeights []int

	var splitNodes []MazeSquare

	// Assigning the first node a weight of 0
	gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = 0

	for !gameGridDFS[int(finishX/20)-1][int(finishY/20)-1].Visited {

		choosingNodes := make(map[MazeSquare]float64)

		// Assigning a new weight to the current node only if it is not the starting point
		if gameGridDFS[int(startX/20)-1][int(startY/20)-1] != gameGridDFS[int(originalStartX/20)-1][int(originalStartY/20)-1] {
			prevWeight += 1
			gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = prevWeight
		}

		// Mark the current node as visited and add the node to the array of nodes for the path taken
		gameGridDFS[int(startX/20)-1][int(startY/20)-1].Visited = true
		bestPath = append(bestPath, gameGridDFS[int(startX/20)-1][int(startY/20)-1])

		// This if block checks if the current node has any neighbours and if so, adds them all sequentially to an array
		// It calculates the distance from the neighbour nodes to the end node
		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasDown && !gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].Visited {
			tempminDistance := euclideanDistance(float64(gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].XCoordinate), float64(gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].YCoordinate), float64(finishX), float64(finishY))
			choosingNodes[gameGridDFS[int(startX/20)-1+1][int(startY/20)-1]] = tempminDistance

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasUp && !gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].Visited {
			tempminDistance := euclideanDistance(float64(gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].XCoordinate), float64(gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].YCoordinate), float64(finishX), float64(finishY))
			choosingNodes[gameGridDFS[int(startX/20)-1-1][int(startY/20)-1]] = tempminDistance

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasLeft && !gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].Visited {
			tempminDistance := euclideanDistance(float64(gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].XCoordinate), float64(gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].YCoordinate), float64(finishX), float64(finishY))
			choosingNodes[gameGridDFS[int(startX/20)-1][int(startY/20)-1-1]] = tempminDistance

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasRight && !gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].Visited {
			tempminDistance := euclideanDistance(float64(gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].XCoordinate), float64(gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].YCoordinate), float64(finishX), float64(finishY))
			choosingNodes[gameGridDFS[int(startX/20)-1][int(startY/20)-1+1]] = tempminDistance

		}

		keys := make([]MazeSquare, 0, len(choosingNodes))

		// The neighbouring nodes are added to a map and then sorted by the distances
		for key := range choosingNodes {
			keys = append(keys, key)
		}

		// This sorts the ( [Node] = Distance ) from highest distance to lowest distance
		sort.SliceStable(keys, func(i, j int) bool {
			return choosingNodes[keys[i]] > choosingNodes[keys[j]]
		})

		// This is adding the sorted nodes back to the array to check for all paths possible
		// This way, the shortest distance nodes are checked first and then the highest distance checked later
		for i := 0; i < len(keys); i++ {
			k := keys[i]
			splitNodes = append(splitNodes, gameGridDFS[int(k.YCoordinate/20)-1][int(k.XCoordinate/20)-1])
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
	fmt.Printf("A* took %s", elapsed)
	fmt.Println("\nA* Concluded\n")
	// Returns the path that the algorithm took to get from the start to the finish
	return bestPath
}

// This function, given the respective x and y values of two nodes, calculates the euclidean distance between them
func euclideanDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	// The euclidean distance is calculated by the square root of the dot product of the difference of the two vectors
	// u = (x1, y1)      v = (x2, y2)     uv = u-v
	// uv . uv = total
	// sqrt(total) = distance

	differenceX := (x2) - (x1)
	differenceY := (y2) - (y1)

	fakeDotProduct := (differenceX * differenceX) + (differenceY * differenceY)

	return math.Sqrt(fakeDotProduct)

}

// This function uses Dijkstras Algorithm to find the shortest path from one node to another in a given maze
func dijkstras(gameGridDFS *[8][8]MazeSquare, startX int, startY int, finishX int, finishY int) []MazeSquare {
	start := time.Now()

	// Storing the original start values
	originalStartX := startX
	originalStartY := startY

	// Marking every node unvisited
	markUnvisited(gameGridDFS)

	var pathTaken []MazeSquare

	prevWeight := 0
	var nodePrevWeights []int

	var splitNodes []MazeSquare

	// Assigning the first node a weight of 0
	gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = 0

	// While the end node of the grid has not been visited

	for !gameGridDFS[int(finishX/20)-1][int(finishY/20)-1].Visited {

		// Assigning a new weight to the current node only if it is not the starting point
		if gameGridDFS[int(startX/20)-1][int(startY/20)-1] != gameGridDFS[(originalStartX/20)-1][(originalStartY/20)-1] {
			prevWeight += 1
			gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = prevWeight
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
	fmt.Println("\nDijkstra Concluded\n")

	// Returns the path that the algorithm took to get from the start to the finish
	return pathTaken
}

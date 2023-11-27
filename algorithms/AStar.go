package algorithms

import (
	"fmt"
	"math"
	"sort"
	"time"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// AStar uses the A* Algorithm to find the shortest path from one node to another in a given maze
// The maze must be built with type mazegrid.Mazesquare
func AStar(gameGridDFS [][]mazegrid.MazeSquare, startX int, startY int, finishX int, finishY int) []mazegrid.MazeSquare {
	start := time.Now() // This is used to time how long the function took to execute

	// Storing the original start values
	originalStartX := startX
	originalStartY := startY

	// Marking every node unvisited
	MarkUnvisited(gameGridDFS)
	// Adding random weights of total weight 100 in the gamegrid
	AddWeights(gameGridDFS, 100)

	var bestPath []mazegrid.MazeSquare // Stores the best path found

	prevWeight := 0           // Stores the previous Node's weight
	var nodePrevWeights []int // Stores the nodes weight while traversing a path

	var splitNodes []mazegrid.MazeSquare

	// Assigning the first node a weight of 0
	gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = 0

	// While the node we want the distance to has not been visited
	for !gameGridDFS[int(finishX/20)-1][int(finishY/20)-1].Visited {

		choosingNodes := make(map[mazegrid.MazeSquare]float64) // Stores all the possible choices that can be made from the current node

		// Assigning a new weight to the current node only if it is not the starting point
		if gameGridDFS[int(startX/20)-1][int(startY/20)-1] != gameGridDFS[int(originalStartX/20)-1][int(originalStartY/20)-1] {
			prevWeight += 1
			gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = prevWeight + gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight

		}

		// Mark the current node as visited and add the node to the array of nodes for the path taken
		gameGridDFS[int(startX/20)-1][int(startY/20)-1].Visited = true
		bestPath = append(bestPath, gameGridDFS[int(startX/20)-1][int(startY/20)-1])

		// This if block checks if the current node has any neighbours and if so, adds them all sequentially to an array
		// It calculates the distance from the neighbour nodes to the end node
		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasWalls.HasDown && !gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].Visited {
			tempminDistance := EuclideanDistance(float64(gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].NodePosition.XCoordinate), float64(gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].NodePosition.YCoordinate), float64(finishX), float64(finishY))
			choosingNodes[gameGridDFS[int(startX/20)-1+1][int(startY/20)-1]] = tempminDistance + float64(gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].Weight)

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasWalls.HasUp && !gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].Visited {
			tempminDistance := EuclideanDistance(float64(gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].NodePosition.XCoordinate), float64(gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].NodePosition.YCoordinate), float64(finishX), float64(finishY))
			choosingNodes[gameGridDFS[int(startX/20)-1-1][int(startY/20)-1]] = tempminDistance + float64(gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].Weight)

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasWalls.HasLeft && !gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].Visited {
			tempminDistance := EuclideanDistance(float64(gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].NodePosition.XCoordinate), float64(gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].NodePosition.YCoordinate), float64(finishX), float64(finishY))
			choosingNodes[gameGridDFS[int(startX/20)-1][int(startY/20)-1-1]] = tempminDistance + float64(gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].Weight)

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasWalls.HasRight && !gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].Visited {
			tempminDistance := EuclideanDistance(float64(gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].NodePosition.XCoordinate), float64(gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].NodePosition.YCoordinate), float64(finishX), float64(finishY))
			choosingNodes[gameGridDFS[int(startX/20)-1][int(startY/20)-1+1]] = tempminDistance + float64(gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].Weight)

		}

		keys := make([]mazegrid.MazeSquare, 0, len(choosingNodes)) // Extracting the keys from the node choices

		// The neighbouring nodes are added to a map based on the keys available
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
			splitNodes = append(splitNodes, gameGridDFS[int(k.NodePosition.YCoordinate/20)-1][int(k.NodePosition.XCoordinate/20)-1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)
		}

		// If no path was possible from the current node, try a previous found neighbour of a node and set that as the new start
		if len(splitNodes) != 0 {
			nodePopped := splitNodes[len(splitNodes)-1]
			splitNodes = splitNodes[:len(splitNodes)-1]

			prevWeight = nodePrevWeights[len(nodePrevWeights)-1]
			nodePrevWeights = nodePrevWeights[:len(nodePrevWeights)-1]

			startY = int(nodePopped.NodePosition.XCoordinate)
			startX = int(nodePopped.NodePosition.YCoordinate)
		}

	}

	elapsed := time.Since(start)
	fmt.Printf("A* took %s", elapsed)
	fmt.Println("\nA* Concluded")
	fmt.Println(" ")
	// Returns the path that the algorithm took to get from the start to the finish
	return bestPath
}

// This function, given the respective x and y values of two nodes, calculates the euclidean distance between them
func EuclideanDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	// The euclidean distance is calculated by the square root of the dot product of the difference of the two vectors
	// u = (x1, y1)      v = (x2, y2)     uv = u-v
	// uv . uv = total
	// sqrt(total) = distance

	differenceX := (x2) - (x1)
	differenceY := (y2) - (y1)

	fakeDotProduct := (differenceX * differenceX) + (differenceY * differenceY)

	return math.Sqrt(fakeDotProduct)

}

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
func AStar(gameGridDFS [][]mazegrid.MazeSquare, startX int, startY int, finishX int, finishY int, squareSize int) []mazegrid.MazeSquare {
	// Marking every node unvisited
	MarkUnvisited(gameGridDFS, false)

	start := time.Now() // This is used to time how long the function took to execute

	// Storing the original start values
	originalStartX := startX
	originalStartY := startY

	startingNode := &gameGridDFS[int(originalStartY/squareSize)-1][int(originalStartX/squareSize)-1]
	endNode := &gameGridDFS[int(finishY/squareSize)-1][int(finishX/squareSize)-1]

	var bestPath []mazegrid.MazeSquare // Stores the best path found

	prevWeight := 0.0             // Stores the previous Node's weight
	var nodePrevWeights []float64 // Stores the nodes weight while traversing a path

	var splitNodes []mazegrid.MazeSquare // Stores the choice to take next when backtracking

	// While the node we want the distance to has not been visited
	for !endNode.Visited {
		fmt.Print("\n")
		currentNode := &gameGridDFS[int(startY/squareSize)-1][int(startX/squareSize)-1]

		choosingNodes := make(map[mazegrid.MazeSquare]float64) // Stores all the possible choices that can be made from the current node

		// Assigning a new weight to the current node only if it is not the starting point
		if currentNode != startingNode {
			prevWeight += 1
			currentNode.Weight = prevWeight + currentNode.Weight

		} else {
			// Assigning the first node a weight of 0
			currentNode.Weight = 0
		}

		// Making sure if for some reason, a node is added that has been visited isn't checked
		if !currentNode.Visited {
			// Mark the current node as visited and add the node to the array of nodes for the path taken
			currentNode.Visited = true
			bestPath = append(bestPath, *currentNode)

			// This if block checks if the current node has any neighbours and if so, adds them all sequentially to an array
			// It calculates the distance from the neighbour nodes to the end node
			if !currentNode.HasWalls.HasDown {
				currentNodeDown := &gameGridDFS[(int(currentNode.Walls.Down.YCoordinate)/squareSize)-1][(int(currentNode.Walls.Down.XCoordinate)/squareSize)-1]

				if !currentNodeDown.Visited {
					tempminDistance := HeuristicsDistance(float64(currentNodeDown.NodePosition.XCoordinate), float64(currentNodeDown.NodePosition.YCoordinate), float64(finishX), float64(finishY))
					choosingNodes[*currentNodeDown] = tempminDistance + float64(currentNodeDown.Weight)
				}

			}

			if !currentNode.HasWalls.HasUp {
				currentNodeUp := &gameGridDFS[(int(currentNode.Walls.Up.YCoordinate)/squareSize)-1][(int(currentNode.Walls.Up.XCoordinate)/squareSize)-1]

				if !currentNodeUp.Visited {
					tempminDistance := HeuristicsDistance(float64(currentNodeUp.NodePosition.XCoordinate), float64(currentNodeUp.NodePosition.YCoordinate), float64(finishX), float64(finishY))
					choosingNodes[*currentNodeUp] = tempminDistance + float64(currentNodeUp.Weight)
				}

			}

			if !currentNode.HasWalls.HasLeft {

				currentNodeLeft := &gameGridDFS[(int(currentNode.Walls.Left.YCoordinate)/squareSize)-1][(int(currentNode.Walls.Left.XCoordinate)/squareSize)-1]

				if !currentNodeLeft.Visited {
					tempminDistance := HeuristicsDistance(float64(currentNodeLeft.NodePosition.XCoordinate), float64(currentNodeLeft.NodePosition.YCoordinate), float64(finishX), float64(finishY))
					choosingNodes[*currentNodeLeft] = tempminDistance + float64(currentNodeLeft.Weight)
				}

			}

			if !currentNode.HasWalls.HasRight {

				currentNodeRight := &gameGridDFS[(int(currentNode.Walls.Right.YCoordinate)/squareSize)-1][(int(currentNode.Walls.Right.XCoordinate)/squareSize)-1]

				if !currentNodeRight.Visited {
					tempminDistance := HeuristicsDistance(float64(currentNodeRight.NodePosition.XCoordinate), float64(currentNodeRight.NodePosition.YCoordinate), float64(finishX), float64(finishY))
					choosingNodes[*currentNodeRight] = tempminDistance + float64(currentNodeRight.Weight)
				}

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
				nodeToAdd := &gameGridDFS[int(k.NodePosition.YCoordinate)/squareSize-1][int(k.NodePosition.XCoordinate)/squareSize-1]
				splitNodes = append(splitNodes, *nodeToAdd)
				nodePrevWeights = append(nodePrevWeights, prevWeight)
			}
		}

		// If no path was possible from the current node, try a previous found neighbour of a node and set that as the new start
		if len(splitNodes) != 0 {
			nodePopped := splitNodes[len(splitNodes)-1]
			splitNodes = splitNodes[:len(splitNodes)-1]

			prevWeight = nodePrevWeights[len(nodePrevWeights)-1]
			nodePrevWeights = nodePrevWeights[:len(nodePrevWeights)-1]

			startX = int(nodePopped.NodePosition.XCoordinate)
			startY = int(nodePopped.NodePosition.YCoordinate)

		}

	}

	elapsed := time.Since(start)
	fmt.Printf("A* took %s", elapsed)
	fmt.Println("\nA* Concluded")
	fmt.Println(" ")

	// Returns the path that the algorithm took to get from the start to the finish
	return bestPath
}

// This function, given the respective x and y values of two nodes, calculates the euclidean distance added to the Manhattan Distance between two points
func HeuristicsDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	// The euclidean distance is calculated by the square root of the dot product of the difference of the two vectors
	// u = (x1, y1)      v = (x2, y2)     uv = u-v
	// uv . uv = total
	// sqrt(total) = distance

	differenceX := (x2) - (x1)
	differenceY := (y2) - (y1)

	// Manhattan is |(X2-X1)| + |(Y2-Y1)|
	manhatten := math.Abs(differenceX) + math.Abs(differenceY)

	fakeDotProduct := (differenceX * differenceX) + (differenceY * differenceY)

	return math.Sqrt(fakeDotProduct) + manhatten

}

func JustPositions(path []mazegrid.MazeSquare) []mazegrid.Position {
	var posArr []mazegrid.Position

	for i := 0; i < len(path); i++ {
		posArr = append(posArr, path[i].NodePosition)
	}

	return posArr
}

// func sliceContains(arrOfSquares []mazegrid.MazeSquare, itemToFind mazegrid.MazeSquare) bool {

// 	for i := 0; i < len(arrOfSquares); i++ {
// 		if arrOfSquares[i] == itemToFind {
// 			return true
// 		}
// 	}

// 	return false
// }

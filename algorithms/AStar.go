package algorithms

import (
	"container/heap"
	"fmt"
	"math"
	"time"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/file"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// AStar uses the A* Algorithm to find the shortest path from one node to another in a given maze
// The maze must be built with type mazegrid.Mazesquare
func AStar(gameGrid [][]mazegrid.MazeSquare, startX, startY, finishX, finishY, squareSize int) []mazegrid.MazeSquare {
	MarkUnvisited(gameGrid, true) // Marking every node unvisited

	start := time.Now()
	fmt.Println("Mem usage before:")
	file.PrintMemUsage()

	// Create a priority queue for choosing next nodes
	priorityQueue := make(mazegrid.PriorityQueue, 0)
	heap.Init(&priorityQueue)

	startNode := &gameGrid[(startY/squareSize)-1][(startX/squareSize)-1]

	endNode := &gameGrid[(finishY/squareSize)-1][(finishX/squareSize)-1]

	// Initialize start node
	startNode.Weight = 0
	startNode.Heuristic = HeuristicsDistance(float64(startX), float64(startY), float64(finishX), float64(finishY))

	pqNode := mazegrid.PriorityNode{}
	pqNode.Init(startNode, startNode.Weight+startNode.Heuristic)

	heap.Push(&priorityQueue, &pqNode)

	// Store each node's predecessor for path reconstruction
	predecessor := make(map[*mazegrid.MazeSquare]*mazegrid.MazeSquare)

	// While the priority queue is not empty
	for len(priorityQueue) > 0 {
		currentNode := heap.Pop(&priorityQueue).(*mazegrid.PriorityNode).GetNode()

		if currentNode == endNode {
			break // Reached the end node
		}

		possibleMoves := getPossibleMoves(gameGrid, currentNode.NodePosition, squareSize)

		for _, move := range possibleMoves {
			nodeToTest := &gameGrid[(int(move.YCoordinate)/squareSize)-1][(int(move.XCoordinate)/squareSize)-1]

			if !nodeToTest.Visited {
				currentNode.Visited = true

				// Update the current distance to the start
				distanceToStart := currentNode.Weight + 1

				if distanceToStart < nodeToTest.Weight {
					predecessor[nodeToTest] = currentNode
					nodeToTest.Weight = distanceToStart
					nodeToTest.Heuristic = HeuristicsDistance(float64(move.XCoordinate), float64(move.YCoordinate), float64(finishX), float64(finishY))

					if !nodeInQueue(nodeToTest, priorityQueue) {
						pqNodeTemp := mazegrid.PriorityNode{}
						pqNodeTemp.Init(nodeToTest, nodeToTest.Weight+nodeToTest.Heuristic)

						//The node will be added to the priority queue, with the both heuristics
						heap.Push(&priorityQueue, &pqNodeTemp)
					}
				}
			}

		}

	}

	// Reconstruct path
	pathTaken := PathReconstructor(startNode, endNode, predecessor)

	elapsed := time.Since(start)
	fmt.Printf("A* took %s", elapsed)
	fmt.Println("\nA* Concluded")
	fmt.Println(" ")

	fmt.Println("Mem usage after:")
	file.PrintMemUsage()
	fmt.Println(" ")

	return pathTaken
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

// This function, given an array of mazeSquares, returns an array of their positions
func JustPositions(path []mazegrid.MazeSquare) []mazegrid.Position {
	var posArr []mazegrid.Position

	for i := 0; i < len(path); i++ {
		posArr = append(posArr, path[i].NodePosition)
	}

	return posArr
}

// This function checks if a node is in the given priority queue.
func nodeInQueue(node *mazegrid.MazeSquare, pq mazegrid.PriorityQueue) bool {
	for _, pn := range pq {
		if pn.GetNode() == node {
			return true
		}
	}
	return false
}

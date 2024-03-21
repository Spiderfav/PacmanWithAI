package algorithms

import (
	"math"
	"math/rand"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

type Algorithm = int

// This enum is used to define the possible algorithms to be used for the ghosts
const (
	DijkstraAlgo   Algorithm = 0
	AStarAlgo      Algorithm = 1
	ReflexAlgo     Algorithm = 2
	MiniMaxAlgo    Algorithm = 3
	BFSAlgo        Algorithm = 4
	DFSAlgo        Algorithm = 5
	ExpectimaxAlgo Algorithm = 6
)

// This simple function is run before any pathfinding algorithm to make sure that the nodes are marked unvisited
// It iterates through a [][]Mazesquare and changes the status of visited, weight and walls to the default.
func MarkUnvisited(gameGrid [][]mazegrid.MazeSquare, markInfinity bool) {

	size := len(gameGrid[0])

	for y := 0; y < size; y++ {

		for x := 0; x < size; x++ {

			gameGrid[y][x].Visited = false
			gameGrid[y][x].Heuristic = 0.0
			gameGrid[y][x].Weight = 0.0

			if markInfinity {
				gameGrid[y][x].Weight = math.Inf(1)
			}

		}

	}

}

// This functions adds weights/pellots to a random square in the grid, given the game grid
// It adds weights/pellots depending on the size of the maze
func AddWeights(gameGrid [][]mazegrid.MazeSquare) {

	numberOfObjects := rand.Intn(len(gameGrid[0]))

	for i := 0; i <= numberOfObjects; i++ {
		xValue := rand.Intn(len(gameGrid[0]) - 1)
		yValue := rand.Intn(len(gameGrid[0]) - 1)

		gameGrid[yValue][xValue].HasPellot = false
		gameGrid[yValue][xValue].HasSuperPellot = true
	}

}

// This function, given a start node, an end node and a dictionary of predecessors, calculates the path taken from start to end
func PathReconstructor(startNode *mazegrid.MazeSquare, endNode *mazegrid.MazeSquare, predecessor map[*mazegrid.MazeSquare]*mazegrid.MazeSquare) []mazegrid.MazeSquare {
	var pathTaken []mazegrid.MazeSquare

	// Start from the end node and work backwards to the start and create the path taken
	for current := endNode; current != nil; current = predecessor[current] {
		pathTaken = append(pathTaken, *current)

		// Break the loop when the start node is reached
		if current == startNode {
			break
		}
	}

	return pathTaken

}

package algorithms

import (
	"math/rand"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

type Algorithm = int

// This enum is used to define the possible algorithms to be used for the ghosts
const (
	DijkstraAlgo Algorithm = 0
	AStarAlgo    Algorithm = 1
	ReflexAlgo   Algorithm = 2
	MiniMaxAlgo  Algorithm = 3
	BFSAlgo      Algorithm = 4
)

// This simple function is run before any pathfinding algorithm to make sure that the nodes are marked unvisited
// It iterates through a [][]Mazesquare and changes the status of visited, weight and walls to the default.
func MarkUnvisited(gameGridDFS [][]mazegrid.MazeSquare) {

	size := len(gameGridDFS[0])

	for y := 0; y < size; y++ {

		for x := 0; x < size; x++ {

			gameGridDFS[y][x].Visited = false

			// If the square contains an object, we need to add a weight to tell the ghosts to try not to traverse through them
			if gameGridDFS[y][x].ContainsObject {
				gameGridDFS[y][x].Weight = 10

			} else {
				gameGridDFS[y][x].Weight = 0

			}

		}

	}

}

// This functions adds weights/pellots to a random square in the grid, given the game grid
// It adds weights/pellots depending on the size of the maze
func AddWeights(gameGridDFS [][]mazegrid.MazeSquare) {

	numberOfObjects := rand.Intn(len(gameGridDFS[0]))

	for i := 0; i <= numberOfObjects; i++ {
		xValue := rand.Intn(len(gameGridDFS[0]) - 1)
		yvalue := rand.Intn(len(gameGridDFS[0]) - 1)

		gameGridDFS[xValue][yvalue].ContainsObject = true
	}

}

package algorithms

import "gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"

// This simple function is run before any pathfinding algorithm to make sure that the nodes are marked unvisited
// It iterates through a [][]Mazesquare and changes the status of visited, weight and walls to the default.
func MarkUnvisited(gameGridDFS [][]mazegrid.MazeSquare) {

	size := len(gameGridDFS[0])

	for y := 0; y < size; y++ {

		for x := 0; x < size; x++ {

			gameGridDFS[y][x].Visited = false
			gameGridDFS[y][x].Weight = 0
			gameGridDFS[y][x].NumberOfWalls = gameGridDFS[y][x].CountWalls()

		}

	}

}

// TODO: Change the function so that a specific X and Y value can be given and weight added to the given MazeSquare

// This functions adds weights to a specific square in the grid
func AddWeights(gameGridDFS [][]mazegrid.MazeSquare, obstacle int) {
	//xValue := rand.Intn(len(gameGridDFS[0]))
	//yvalue := rand.Intn(len(gameGridDFS[0]))

	gameGridDFS[0][1].Weight = obstacle
}

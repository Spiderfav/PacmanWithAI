package algorithms

import "gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"

// This simple function is run before any pathfinding algorithm to make sure that the nodes are marked unvisited
func MarkUnvisited(gameGridDFS [][]mazegrid.MazeSquare) {

	size := len(gameGridDFS[0])

	for y := 0; y < size; y++ {

		for x := 0; x < size; x++ {

			gameGridDFS[y][x].Visited = false
			gameGridDFS[y][x].Weight = 0
			gameGridDFS[y][x].NumberOfWalls = mazegrid.CountWalls(gameGridDFS[y][x])

		}

	}

}

// This functions adds weights to a specific square in the grid
func AddWeights(gameGridDFS [][]mazegrid.MazeSquare, obstacle int) {
	//xValue := rand.Intn(len(gameGridDFS[0]))
	//yvalue := rand.Intn(len(gameGridDFS[0]))

	gameGridDFS[0][1].Weight = obstacle
}

package algorithms

import (
	"math"
	"math/rand"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// This function returns the path for a Reflex agent to taken, given the position of the player and the pellots on the game board
func Reflex(gameGrid [][]mazegrid.MazeSquare, playerPos mazegrid.Position, ghostPos mazegrid.Position, pellots []mazegrid.Position, squareSize int, algo Algorithm) []mazegrid.MazeSquare {

	//Check to see how many pellots are left, if less than 8, go to Pacman
	if len(pellots) <= 8 {
		return pathtoEnd(gameGrid, ghostPos, playerPos, squareSize, algo)
	}

	distance := HeuristicsDistance(float64(ghostPos.XCoordinate), float64(ghostPos.YCoordinate), float64(playerPos.XCoordinate), float64(playerPos.YCoordinate))

	// If the player is close, go to the player
	if distance <= float64(squareSize)*8 {
		return pathtoEnd(gameGrid, ghostPos, playerPos, squareSize, algo)
	}

	nearestPellotPos, nearestPellotDistance := nearestPellot(ghostPos, pellots)

	// If the ghost is near a pellot, go to the pellot
	if nearestPellotDistance >= float64(squareSize) && nearestPellotDistance <= float64(squareSize)*8 {

		return pathtoEnd(gameGrid, ghostPos, nearestPellotPos, squareSize, algo)

	}

	// Else go to a random place in the game grid
	randomX := rand.Intn(len(gameGrid) - 1)
	randomY := rand.Intn(len(gameGrid) - 1)

	randomNode := gameGrid[randomY][randomX].NodePosition

	return pathtoEnd(gameGrid, ghostPos, randomNode, squareSize, algo)

}

// This function, given a character and a array of pellots, will return the position and the distance to the nearest pellot
func nearestPellot(characterPosition mazegrid.Position, pellots []mazegrid.Position) (mazegrid.Position, float64) {

	nearestPellotPos := mazegrid.Position{}
	nearestPellotDistance := math.Inf(1)

	for i := 0; i < len(pellots); i++ {

		distance := HeuristicsDistance(float64(characterPosition.XCoordinate), float64(characterPosition.YCoordinate), float64(pellots[i].XCoordinate), float64(pellots[i].YCoordinate))

		if distance < nearestPellotDistance {
			nearestPellotPos = pellots[i]
			nearestPellotDistance = distance
		}
	}

	return nearestPellotPos, nearestPellotDistance

}

// This function, given the algorithm the ghost is using, returns the path to take to the end goal
func pathtoEnd(gameGrid [][]mazegrid.MazeSquare, ghostPos mazegrid.Position, endPos mazegrid.Position, squareSize int, algo Algorithm) []mazegrid.MazeSquare {
	var path []mazegrid.MazeSquare

	switch algo {
	case DijkstraAlgo:
		path = Dijkstras(gameGrid, int(ghostPos.XCoordinate), int(ghostPos.YCoordinate), int(endPos.XCoordinate), int(endPos.YCoordinate), squareSize)

	case AStarAlgo:
		path = AStar(gameGrid, int(ghostPos.XCoordinate), int(ghostPos.YCoordinate), int(endPos.XCoordinate), int(endPos.YCoordinate), squareSize)

	case BFSAlgo:
		path = BFS(gameGrid, int(ghostPos.XCoordinate), int(ghostPos.YCoordinate), int(endPos.XCoordinate), int(endPos.YCoordinate), squareSize)

	case DFSAlgo:
		path, _ = AbsolutePath(DFSearch(gameGrid, int(ghostPos.XCoordinate), int(ghostPos.YCoordinate), int(endPos.XCoordinate), int(endPos.YCoordinate), squareSize))

	}

	return path
}

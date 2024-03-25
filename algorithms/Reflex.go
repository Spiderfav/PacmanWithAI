package algorithms

import (
	"math"
	"math/rand"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// This function returns the path for a Reflex agent to taken, given the position of the player and the pellots on the game board
func Reflex(gameGridDFS [][]mazegrid.MazeSquare, playerPos mazegrid.Position, ghostPos mazegrid.Position, pellots []mazegrid.Position, squareSize int) []mazegrid.MazeSquare {

	//Check to see how many pellots are left, if less than 8, go to Pacman
	if len(pellots) <= 8 {
		return AStar(gameGridDFS, int(ghostPos.XCoordinate), int(ghostPos.YCoordinate), int(playerPos.XCoordinate), int(playerPos.YCoordinate), squareSize)
	}

	distance := HeuristicsDistance(float64(ghostPos.XCoordinate), float64(ghostPos.YCoordinate), float64(playerPos.XCoordinate), float64(playerPos.YCoordinate))

	// If the player is close, go to the player
	if distance <= 160 {

		return AStar(gameGridDFS, int(ghostPos.XCoordinate), int(ghostPos.YCoordinate), int(playerPos.XCoordinate), int(playerPos.YCoordinate), squareSize)
	}

	nearestPellotPos, nearestPellotDistance := nearestPellot(ghostPos, pellots)

	// If the ghost is near a pellot, go to the pellot
	if nearestPellotDistance >= 20 && nearestPellotDistance <= 160 {

		return AStar(gameGridDFS, int(ghostPos.XCoordinate), int(ghostPos.YCoordinate), int(nearestPellotPos.XCoordinate), int(nearestPellotPos.YCoordinate), squareSize)

	}

	// Else go to a random place in the game grid
	randomX := rand.Intn(len(gameGridDFS) - 1)
	randomY := rand.Intn(len(gameGridDFS) - 1)

	randomNode := gameGridDFS[randomY][randomX].NodePosition

	return AStar(gameGridDFS, int(ghostPos.XCoordinate), int(ghostPos.YCoordinate), int(randomNode.XCoordinate), int(randomNode.YCoordinate), squareSize)

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

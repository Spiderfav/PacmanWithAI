package algorithms

import (
	"math"
	"math/rand"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

func Reflex(gameGridDFS [][]mazegrid.MazeSquare, playerPos mazegrid.Position, ghostPos mazegrid.Position, pellots []mazegrid.Position) []mazegrid.MazeSquare {

	//Check for distance to Player

	distance := HeuristicsDistance(float64(ghostPos.XCoordinate), float64(ghostPos.YCoordinate), float64(playerPos.XCoordinate), float64(playerPos.YCoordinate))

	if distance <= 60 {

		return AStar(gameGridDFS, int(ghostPos.YCoordinate), int(ghostPos.XCoordinate), int(playerPos.YCoordinate), int(playerPos.XCoordinate))
	}

	nearestPellotPos, nearestPellotDistance := nearestPellot(ghostPos, pellots)

	if nearestPellotDistance >= 20 && nearestPellotDistance <= 100 {

		return AStar(gameGridDFS, int(ghostPos.YCoordinate), int(ghostPos.XCoordinate), int(nearestPellotPos.YCoordinate), int(nearestPellotPos.XCoordinate))

	}

	randomX := rand.Intn(len(gameGridDFS) - 1)
	randomY := rand.Intn(len(gameGridDFS) - 1)

	randomNode := gameGridDFS[randomX][randomY].NodePosition

	return AStar(gameGridDFS, int(ghostPos.YCoordinate), int(ghostPos.XCoordinate), int(randomNode.YCoordinate), int(randomNode.XCoordinate))

}

func nearestPellot(ghostPos mazegrid.Position, pellots []mazegrid.Position) (mazegrid.Position, float64) {

	nearestPellotPos := mazegrid.Position{}
	nearestPellotDistance := math.Inf(1)

	for i := 0; i < len(pellots); i++ {

		distance := HeuristicsDistance(float64(ghostPos.XCoordinate), float64(ghostPos.YCoordinate), float64(pellots[i].XCoordinate), float64(pellots[i].YCoordinate))

		if distance < nearestPellotDistance {
			nearestPellotPos = pellots[i]
			nearestPellotDistance = distance
		}
	}

	return nearestPellotPos, nearestPellotDistance

}

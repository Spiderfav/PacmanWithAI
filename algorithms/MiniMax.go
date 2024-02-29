package algorithms

import (
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

func MiniMax() {}

func evalPos(pacmanPos mazegrid.Position, pacmanPoints int, ghostPos mazegrid.Position, pellots []mazegrid.Position, isGhost bool) int {
	totalEval := 0
	nearestPellotPos, nearestPellotDistance := nearestPellot(pacmanPos, pellots)

	if isGhost {
		totalEval = totalEval + len(pellots)
		distanceToNearestPacmanPellot := HeuristicsDistance(float64(ghostPos.XCoordinate), float64(ghostPos.YCoordinate), float64(nearestPellotPos.XCoordinate), float64(nearestPellotPos.YCoordinate))
		totalEval = int((float64(totalEval) + distanceToNearestPacmanPellot) * -1)

	} else {
		totalEval = totalEval + pacmanPoints
		totalEval = totalEval + int(nearestPellotDistance)
	}

	return totalEval
}

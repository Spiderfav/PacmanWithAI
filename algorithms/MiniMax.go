package algorithms

import (
	"math"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

type PruningParams struct {
	Alpha float64
	Beta  float64
}

func MiniMax(gameGrid [][]mazegrid.MazeSquare, params PruningParams, pacmanPos []mazegrid.Position, pacmanPoints int, ghostPos []mazegrid.Position, pellots []mazegrid.Position, depthToSearch int, isPacman bool, usePruning bool) (int, []mazegrid.Position, []mazegrid.Position, PruningParams) {

	if depthToSearch == 0 || pacmanPos[len(pacmanPos)-1] == ghostPos[len(ghostPos)-1] {
		return evalPos(pacmanPos[len(pacmanPos)-1], pacmanPoints, ghostPos[len(ghostPos)-1], pellots, isPacman), pacmanPos, ghostPos, params
	}

	if isPacman {
		maxEval := math.Inf(-1)
		var bestPacmanPos []mazegrid.Position

		possibleMoves := getPossibleMoves(gameGrid, pacmanPos[len(pacmanPos)-1])

		for _, element := range possibleMoves {
			tempPacmanPos := make([]mazegrid.Position, len(pacmanPos))
			copy(tempPacmanPos, pacmanPos)
			tempPacmanPos = append(tempPacmanPos, element)

			eval, newPacmanPos, newGhostPos, _ := MiniMax(gameGrid, params, tempPacmanPos, pacmanPoints, ghostPos, pellots, depthToSearch-1, false, usePruning)

			if float64(eval) > maxEval {
				maxEval = float64(eval)
				bestPacmanPos = newPacmanPos
				ghostPos = newGhostPos
			}

			if usePruning {
				if float64(eval) > params.Alpha {
					params.Alpha = float64(eval)
				}

				if params.Beta <= params.Alpha {
					break
				}
			}

		}

		return int(maxEval), bestPacmanPos, ghostPos, params

	} else {
		minEval := math.Inf(1)
		var bestGhostPos []mazegrid.Position

		possibleMoves := getPossibleMoves(gameGrid, ghostPos[len(ghostPos)-1])

		for _, element := range possibleMoves {
			tempGhostPos := make([]mazegrid.Position, len(ghostPos))
			copy(tempGhostPos, ghostPos)
			tempGhostPos = append(tempGhostPos, element)

			eval, newPacmanPos, newGhostPos, _ := MiniMax(gameGrid, params, pacmanPos, pacmanPoints, tempGhostPos, pellots, depthToSearch-1, true, usePruning)

			if float64(eval) < minEval {
				minEval = float64(eval)
				pacmanPos = newPacmanPos
				bestGhostPos = newGhostPos

			}

			if usePruning {
				if float64(eval) < params.Alpha {
					params.Alpha = float64(eval)
				}

				if params.Alpha <= params.Beta {
					break
				}
			}

		}

		return int(minEval), pacmanPos, bestGhostPos, params
	}
}

func evalPos(pacmanPos mazegrid.Position, pacmanPoints int, ghostPos mazegrid.Position, pellots []mazegrid.Position, isPacman bool) int {
	totalEval := 0
	nearestPellotPos, nearestPellotDistance := nearestPellot(pacmanPos, pellots)

	if !isPacman {
		totalEval = totalEval + len(pellots)
		totalEval = int((float64(totalEval) + nearestPellotDistance) * -1)

	} else {
		totalEval = totalEval + pacmanPoints
		distanceToNearestPacmanPellot := HeuristicsDistance(float64(ghostPos.XCoordinate), float64(ghostPos.YCoordinate), float64(nearestPellotPos.XCoordinate), float64(nearestPellotPos.YCoordinate))
		totalEval = totalEval + int(distanceToNearestPacmanPellot)
	}

	return totalEval
}

func getPossibleMoves(gameGrid [][]mazegrid.MazeSquare, charPos mazegrid.Position) []mazegrid.Position {
	// Try up, down, left, right
	var possibleMoves []mazegrid.Position

	firstArr := int((charPos.YCoordinate / 20) - 1)
	secondArr := int((charPos.XCoordinate / 20) - 1)

	gameNode := gameGrid[firstArr][secondArr]

	if !gameNode.HasWalls.HasUp {
		possibleMoves = append(possibleMoves, gameNode.Walls.Up)
	}
	if !gameNode.HasWalls.HasDown {
		possibleMoves = append(possibleMoves, gameNode.Walls.Down)
	}
	if !gameNode.HasWalls.HasLeft {
		possibleMoves = append(possibleMoves, gameNode.Walls.Left)
	}
	if !gameNode.HasWalls.HasRight {
		possibleMoves = append(possibleMoves, gameNode.Walls.Right)
	}

	return possibleMoves
}

func ReversePath(s []mazegrid.MazeSquare) []mazegrid.MazeSquare {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

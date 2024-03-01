package algorithms

import (
	"fmt"
	"math"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

func MiniMax(gameGrid [][]mazegrid.MazeSquare, pacmanPos []mazegrid.Position, pacmanPoints int, ghostPos []mazegrid.Position, pellots []mazegrid.Position, depthToSearch int, isPacman bool) (int, []mazegrid.Position, []mazegrid.Position) {

	fmt.Println("Pacman Inside: ", pacmanPos)
	fmt.Println("Ghost Inside: ", ghostPos)

	if depthToSearch == 0 || pacmanPos[len(pacmanPos)-1] == ghostPos[len(ghostPos)-1] {
		return evalPos(pacmanPos[len(pacmanPos)-1], pacmanPoints, ghostPos[len(ghostPos)-1], pellots, isPacman), pacmanPos, ghostPos
	}

	if isPacman {
		maxEval := math.Inf(-1)
		var bestPacmanPos []mazegrid.Position

		possibleMoves := getPossibleMoves(gameGrid, pacmanPos[len(pacmanPos)-1])

		for _, element := range possibleMoves {
			tempPacmanPos := make([]mazegrid.Position, len(pacmanPos))
			copy(tempPacmanPos, pacmanPos)
			tempPacmanPos = append(tempPacmanPos, element)

			eval, newPacmanPos, newGhostPos := MiniMax(gameGrid, tempPacmanPos, pacmanPoints, ghostPos, pellots, depthToSearch-1, false)

			if float64(eval) > maxEval {
				maxEval = float64(eval)
				bestPacmanPos = newPacmanPos
				ghostPos = newGhostPos
			}
		}

		return int(maxEval), bestPacmanPos, ghostPos

	} else {
		minEval := math.Inf(1)
		var bestGhostPos []mazegrid.Position

		possibleMoves := getPossibleMoves(gameGrid, ghostPos[len(ghostPos)-1])

		for _, element := range possibleMoves {
			tempGhostPos := make([]mazegrid.Position, len(ghostPos))
			copy(tempGhostPos, ghostPos)
			tempGhostPos = append(tempGhostPos, element)

			eval, newPacmanPos, newGhostPos := MiniMax(gameGrid, pacmanPos, pacmanPoints, tempGhostPos, pellots, depthToSearch-1, true)

			if float64(eval) < minEval {
				minEval = float64(eval)
				pacmanPos = newPacmanPos
				bestGhostPos = newGhostPos
			}
		}

		return int(minEval), pacmanPos, bestGhostPos
	}
}

func evalPos(pacmanPos mazegrid.Position, pacmanPoints int, ghostPos mazegrid.Position, pellots []mazegrid.Position, isPacman bool) int {
	totalEval := 0
	nearestPellotPos, nearestPellotDistance := nearestPellot(pacmanPos, pellots)

	if !isPacman {
		totalEval = totalEval + len(pellots)
		distanceToNearestPacmanPellot := HeuristicsDistance(float64(ghostPos.XCoordinate), float64(ghostPos.YCoordinate), float64(nearestPellotPos.XCoordinate), float64(nearestPellotPos.YCoordinate))
		totalEval = int((float64(totalEval) + distanceToNearestPacmanPellot) * -1)

	} else {
		totalEval = totalEval + pacmanPoints
		totalEval = totalEval + int(nearestPellotDistance)
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

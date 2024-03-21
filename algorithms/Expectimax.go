package algorithms

import (
	"math"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

func Expectimax(gameGrid [][]mazegrid.MazeSquare, pacmanPos []mazegrid.Position, pacmanPoints int, ghostPos []mazegrid.Position, pellots []mazegrid.Position, depthToSearch int, isPacman bool, squareSize int) (int, []mazegrid.Position, []mazegrid.Position) {

	if depthToSearch == 0 || pacmanPos[len(pacmanPos)-1] == ghostPos[len(ghostPos)-1] {
		return evalPos(pacmanPos[len(pacmanPos)-1], pacmanPoints, ghostPos[len(ghostPos)-1], pellots, isPacman), pacmanPos, ghostPos
	}

	if isPacman {
		maxEval := math.Inf(-1)
		var bestPacmanPos []mazegrid.Position

		possibleMoves := getPossibleMoves(gameGrid, pacmanPos[len(pacmanPos)-1], squareSize)

		for _, element := range possibleMoves {
			tempPacmanPos := make([]mazegrid.Position, len(pacmanPos))
			copy(tempPacmanPos, pacmanPos)
			tempPacmanPos = append(tempPacmanPos, element)

			eval, newPacmanPos, newGhostPos := Expectimax(gameGrid, tempPacmanPos, pacmanPoints, ghostPos, pellots, depthToSearch-1, false, squareSize)

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

		possibleMoves := getPossibleMoves(gameGrid, ghostPos[len(ghostPos)-1], squareSize)

		for _, element := range possibleMoves {
			tempGhostPos := make([]mazegrid.Position, len(ghostPos))
			copy(tempGhostPos, ghostPos)
			tempGhostPos = append(tempGhostPos, element)

			eval, newPacmanPos, newGhostPos := Expectimax(gameGrid, pacmanPos, pacmanPoints, tempGhostPos, pellots, depthToSearch-1, true, squareSize)

			if float64(eval) < minEval {
				minEval = float64(eval)
				pacmanPos = newPacmanPos
				bestGhostPos = newGhostPos
			}
		}

		return int(minEval), pacmanPos, bestGhostPos
	}
}

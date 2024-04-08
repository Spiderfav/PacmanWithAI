package algorithms

import (
	"math"
	"math/rand"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// In Params: Game Grid, Moves Pacman will make, Current Pacman points, Moves Ghosts will make, Position of all the pellots in the maze, Max tree depth, If Maximiser or Minimiser.
// Out Params: Evaluation of Position, Moves Pacman made to given position eval, Moves Ghost made to given position eval.
// This function will calculate play between two opponents, one will be maximising whereas the other will be randomised, given a maze and the pellots in the maze.
func Expectimax(gameGrid [][]mazegrid.MazeSquare, pacmanPos []mazegrid.Position, pacmanPoints int, ghostPos []mazegrid.Position, pellots []mazegrid.Position, depthToSearch int, isPacman bool, squareSize int) (int, []mazegrid.Position, []mazegrid.Position) {

	// Pacman is always assumed to be the Maximiser

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
		var totalEval float64
		var bestGhostPos []mazegrid.Position

		possibleMoves := getPossibleMoves(gameGrid, ghostPos[len(ghostPos)-1], squareSize)

		randomNodeChosen := rand.Intn(len(possibleMoves) - 1)

		for i, element := range possibleMoves {
			tempGhostPos := make([]mazegrid.Position, len(ghostPos))
			copy(tempGhostPos, ghostPos)
			tempGhostPos = append(tempGhostPos, element)

			eval, newPacmanPos, newGhostPos := Expectimax(gameGrid, pacmanPos, pacmanPoints, tempGhostPos, pellots, depthToSearch-1, true, squareSize)
			totalEval = totalEval + float64(eval)

			if i == randomNodeChosen {
				pacmanPos = newPacmanPos
				bestGhostPos = newGhostPos
			}

		}

		avgEval := totalEval / float64(len(possibleMoves))
		return int(avgEval), pacmanPos, bestGhostPos
	}
}

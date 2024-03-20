package algorithms

import (
	"math"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// These are the parameters to be used to prune branches in perfect play
type PruningParams struct {
	Alpha float64
	Beta  float64
}

// In Params: Game Grid, Pruning Paramaters, Moves Pacman will make, Current Pacman points, Moves Ghosts will make, Position of all the pellots in the maze, Max tree depth, If Maximiser or Minimiser, Whether to use pruning or not
// Out Params: Evaluation of Position, Moves Pacman made to given position eval, Moves Ghost made to given position eval
// This functions will calculate best play between two opponents, given a maze and the pellots in the maze. It can allow play to be pruned or not
func MiniMax(gameGrid [][]mazegrid.MazeSquare, params PruningParams, pacmanPos []mazegrid.Position, pacmanPoints int, ghostPos []mazegrid.Position, pellots []mazegrid.Position, depthToSearch int, isPacman bool, usePruning bool, squareSize int) (int, []mazegrid.Position, []mazegrid.Position, PruningParams) {

	// If the depth has reached zero or ghost has caught pacman
	if depthToSearch == 0 || pacmanPos[len(pacmanPos)-1] == ghostPos[len(ghostPos)-1] {
		return evalPos(pacmanPos[len(pacmanPos)-1], pacmanPoints, ghostPos[len(ghostPos)-1], pellots, isPacman), pacmanPos, ghostPos, params
	}

	if isPacman {
		maxEval := math.Inf(-1)
		var bestPacmanPos []mazegrid.Position

		// Get the possible moves from this single square
		possibleMoves := getPossibleMoves(gameGrid, pacmanPos[len(pacmanPos)-1], squareSize)

		// For every possible move in the square
		for _, element := range possibleMoves {
			tempPacmanPos := make([]mazegrid.Position, len(pacmanPos))
			copy(tempPacmanPos, pacmanPos)
			tempPacmanPos = append(tempPacmanPos, element)

			eval, newPacmanPos, newGhostPos, _ := MiniMax(gameGrid, params, tempPacmanPos, pacmanPoints, ghostPos, pellots, depthToSearch-1, false, usePruning, squareSize)

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

		possibleMoves := getPossibleMoves(gameGrid, ghostPos[len(ghostPos)-1], squareSize)

		for _, element := range possibleMoves {
			tempGhostPos := make([]mazegrid.Position, len(ghostPos))
			copy(tempGhostPos, ghostPos)
			tempGhostPos = append(tempGhostPos, element)

			eval, newPacmanPos, newGhostPos, _ := MiniMax(gameGrid, params, pacmanPos, pacmanPoints, tempGhostPos, pellots, depthToSearch-1, true, usePruning, squareSize)

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

// This function, given the position of the player and ghost and the number of pellots left will evaluate how good a given move is
func evalPos(pacmanPos mazegrid.Position, pacmanPoints int, ghostPos mazegrid.Position, pellots []mazegrid.Position, isPacman bool) int {
	totalEval := 0

	// Get the nearest pellot to pacman
	nearestPellotPos, nearestPellotDistance := nearestPellot(pacmanPos, pellots)

	// Get the distance between Pacman and the Ghost
	distanceBetween := HeuristicsDistance(float64(ghostPos.XCoordinate), float64(ghostPos.YCoordinate), float64(pacmanPos.XCoordinate), float64(pacmanPos.YCoordinate))

	// If it is the ghosts evaluation
	if !isPacman {

		// Check how many pellots are left
		totalEval = totalEval + len(pellots)

		// Final evaluation is the sum of the pellots left plus how far away pacman is plus the distance between the ghost and pacman
		// The distance between the ghost and pacman is summed here to encourage the ghost to chase
		totalEval = int((float64(totalEval)+nearestPellotDistance)*-1) + int(distanceBetween)

	} else {

		// Check how many pellots pacman collected
		totalEval = totalEval + pacmanPoints

		// Get the distance frrom the ghost to the nearest pellot
		distanceToNearestPacmanPellot := HeuristicsDistance(float64(ghostPos.XCoordinate), float64(ghostPos.YCoordinate), float64(nearestPellotPos.XCoordinate), float64(nearestPellotPos.YCoordinate))

		// Final evaluation is the sum of the pellots collected plus how far away pacman from the nearest pellot plus the distance between the ghost and pacman
		totalEval = totalEval + int(distanceToNearestPacmanPellot) + int(distanceBetween)

	}

	return totalEval
}

// This function, given the game grid and a position, will look at the possible moves in that given square
func getPossibleMoves(gameGrid [][]mazegrid.MazeSquare, charPos mazegrid.Position, squareSize int) []mazegrid.Position {
	// Try up, down, left, right
	var possibleMoves []mazegrid.Position

	firstArr := int((charPos.YCoordinate / float32(squareSize)) - 1)
	secondArr := int((charPos.XCoordinate / float32(squareSize)) - 1)

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

// This function, given an array of maze squares, will reverse the order for traversal
func ReversePath(s []mazegrid.MazeSquare) []mazegrid.MazeSquare {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

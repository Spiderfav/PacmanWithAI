package algorithms

import "gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"

// This function, given a array of positions and a game grid, turns the array of positions to an array of maze squares
func PosToNode(gameGrid [][]mazegrid.MazeSquare, arrOfPos []mazegrid.Position, squareSize int) []mazegrid.MazeSquare {
	var posToNodeArr []mazegrid.MazeSquare

	for i := 0; i < len(arrOfPos); i++ {
		firstArr := int((int(arrOfPos[i].YCoordinate) / squareSize) - 1)
		secondArr := int((int(arrOfPos[i].XCoordinate) / squareSize) - 1)

		posToNodeArr = append(posToNodeArr, gameGrid[firstArr][secondArr])

	}

	return posToNodeArr

}

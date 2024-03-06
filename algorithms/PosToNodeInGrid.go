package algorithms

import "gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"

func PosToNode(gameGrid [][]mazegrid.MazeSquare, arrOfPos []mazegrid.Position) []mazegrid.MazeSquare {
	var posToNodeArr []mazegrid.MazeSquare

	for i := 0; i < len(arrOfPos); i++ {
		firstArr := int((arrOfPos[i].YCoordinate / 20) - 1)
		secondArr := int((arrOfPos[i].XCoordinate / 20) - 1)

		posToNodeArr = append(posToNodeArr, gameGrid[firstArr][secondArr])

	}

	return posToNodeArr

}

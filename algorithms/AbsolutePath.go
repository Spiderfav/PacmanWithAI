package algorithms

import (
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// This function, given the path made from one of the maze searching algorithms, provides the final path from start to finish
func AbsolutePath(pathTaken []mazegrid.MazeSquare) ([]mazegrid.MazeSquare, int) {
	var finalPath []mazegrid.MazeSquare
	var totalWeight int

	currentNode := pathTaken[len(pathTaken)-1]
	finalPath = append(finalPath, currentNode)
	totalWeight = totalWeight + currentNode.Weight

	//From the end node to the start, it finds the path the Single-Agent should take
	// At a given node, it looks at it's neighbours and looks in the next node in the pathTaken and if it is a neighbour then it's added to the finalPath
	for i := len(pathTaken) - 1; i > 0; i-- {

		if i == 0 {
			currentNode = pathTaken[i-1]
			finalPath = append(finalPath, pathTaken[i])
			continue
		}

		if !currentNode.HasLeft {
			if currentNode.Left.XCoordinate == pathTaken[i-1].XCoordinate && currentNode.Left.YCoordinate == pathTaken[i-1].YCoordinate {
				currentNode = pathTaken[i-1]
				finalPath = append(finalPath, pathTaken[i-1])
				totalWeight = totalWeight + currentNode.Weight
				continue
			}
		}
		if !currentNode.HasRight {
			if currentNode.Right.XCoordinate == pathTaken[i-1].XCoordinate && currentNode.Right.YCoordinate == pathTaken[i-1].YCoordinate {
				currentNode = pathTaken[i-1]
				finalPath = append(finalPath, pathTaken[i-1])
				totalWeight = totalWeight + currentNode.Weight
				continue
			}
		}

		if !currentNode.HasDown {
			if currentNode.Down.XCoordinate == pathTaken[i-1].XCoordinate && currentNode.Down.YCoordinate == pathTaken[i-1].YCoordinate {
				currentNode = pathTaken[i-1]
				finalPath = append(finalPath, pathTaken[i-1])
				totalWeight = totalWeight + currentNode.Weight
				continue
			}

		}

		if !currentNode.HasUp {
			if currentNode.Up.XCoordinate == pathTaken[i-1].XCoordinate && currentNode.Up.YCoordinate == pathTaken[i-1].YCoordinate {
				currentNode = pathTaken[i-1]
				finalPath = append(finalPath, pathTaken[i-1])
				totalWeight = totalWeight + currentNode.Weight
				continue
			}

		}

		continue

	}

	return finalPath, totalWeight
}

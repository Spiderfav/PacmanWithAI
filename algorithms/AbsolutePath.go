package algorithms

import (
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// AbsolutePath takes a path generated by a maze searching algorithm
// and returns the final path from start to finish along with its total weight
func AbsolutePath(pathTaken []mazegrid.MazeSquare) ([]mazegrid.MazeSquare, int) {
	var finalPath []mazegrid.MazeSquare // Stores the final path
	var totalWeight int                 // Accumulates the total weight of the path

	// Start with the last node in the given path
	currentNode := pathTaken[len(pathTaken)-1]
	finalPath = append(finalPath, currentNode)
	totalWeight += currentNode.Weight // Add the weight of the last node to the total

	// Iterate through the path in reverse to reconstruct the absolute path
	for i := len(pathTaken) - 1; i > 0; i-- {
		if i == 0 {
			// If reached the start node, add it to the final path
			currentNode = pathTaken[i-1]
			finalPath = append(finalPath, pathTaken[i])
			continue
		}

		// Check each direction of the current node to identify the next node in the path
		// If found, update the current node and add it to the final path while updating the total weight
		if !currentNode.HasLeft {
			if currentNode.Left.XCoordinate == pathTaken[i-1].XCoordinate && currentNode.Left.YCoordinate == pathTaken[i-1].YCoordinate {
				currentNode, finalPath, totalWeight = appendNode(i-1, pathTaken, finalPath, totalWeight)
				continue
			}
		}

		// Similar checks are performed for other directions (Right, Down, Up)

		if !currentNode.HasRight {
			if currentNode.Right.XCoordinate == pathTaken[i-1].XCoordinate && currentNode.Right.YCoordinate == pathTaken[i-1].YCoordinate {
				currentNode, finalPath, totalWeight = appendNode(i-1, pathTaken, finalPath, totalWeight)
				continue
			}
		}

		if !currentNode.HasDown {
			if currentNode.Down.XCoordinate == pathTaken[i-1].XCoordinate && currentNode.Down.YCoordinate == pathTaken[i-1].YCoordinate {
				currentNode, finalPath, totalWeight = appendNode(i-1, pathTaken, finalPath, totalWeight)
				continue
			}

		}

		if !currentNode.HasUp {
			if currentNode.Up.XCoordinate == pathTaken[i-1].XCoordinate && currentNode.Up.YCoordinate == pathTaken[i-1].YCoordinate {
				currentNode, finalPath, totalWeight = appendNode(i-1, pathTaken, finalPath, totalWeight)
				continue
			}

		}
		// Continue the loop if the next node is not found in any direction
		continue
	}

	return finalPath, totalWeight // Return the final path and its total weight
}

// appendNode takes the position of the node, the path taken and the final path
// It returns the node to take next, the updated path and the weight of the node
func appendNode(i int, pathTaken []mazegrid.MazeSquare, finalPath []mazegrid.MazeSquare, totalWeight int) (mazegrid.MazeSquare, []mazegrid.MazeSquare, int) {
	currentNode := pathTaken[i]
	finalPath = append(finalPath, pathTaken[i])
	totalWeight = totalWeight + currentNode.Weight

	return currentNode, finalPath, totalWeight
}

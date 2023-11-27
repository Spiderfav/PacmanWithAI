package generation

import (
	"fmt"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

type TreeNode struct {
	CurrentNode mazegrid.MazeSquare
	LeftNode    *mazegrid.MazeSquare
	RightNode   *mazegrid.MazeSquare
}

func MazeToGraph(gameGridDFS [][]mazegrid.MazeSquare, startX float32, startY float32, endX float32, endY float32) []mazegrid.MazeSquare {
	algorithms.MarkUnvisited(gameGridDFS)
	var definiteNodes []mazegrid.MazeSquare

	size := len(gameGridDFS[0])

	for y := 0; y < size; y++ {

		for x := 0; x < size; x++ {

			if (gameGridDFS[y][x].NodePosition.XCoordinate == startX) && (gameGridDFS[y][x].NodePosition.YCoordinate == startY) || (gameGridDFS[y][x].NodePosition.XCoordinate == endX) && (gameGridDFS[y][x].NodePosition.YCoordinate == endY) {
				definiteNodes = append(definiteNodes, gameGridDFS[y][x])
				continue
			}

			if gameGridDFS[y][x].NumberOfWalls == 3 || gameGridDFS[y][x].NumberOfWalls == 1 {
				definiteNodes = append(definiteNodes, gameGridDFS[y][x])
			}

		}

	}

	return definiteNodes
}

func AllPaths(gameGridDFS [][]mazegrid.MazeSquare, definiteNodes []mazegrid.MazeSquare) [][]mazegrid.MazeSquare {
	var paths [][]mazegrid.MazeSquare

	for i := 1; i < len(definiteNodes); i++ {
		fmt.Println("Creating Graph")
		pathTaken := algorithms.Dijkstras(gameGridDFS, int(definiteNodes[0].NodePosition.XCoordinate), int(definiteNodes[0].NodePosition.YCoordinate), int(definiteNodes[i].NodePosition.XCoordinate), int(definiteNodes[i].NodePosition.YCoordinate))
		finalPath, _ := algorithms.AbsolutePath(pathTaken)
		paths = append(paths, finalPath)
	}

	return paths
}

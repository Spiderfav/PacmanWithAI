package main

type TreeNode struct {
	CurrentNode MazeSquare
	LeftNode    *MazeSquare
	RightNode   *MazeSquare
}

func mazeToGraph(gameGridDFS [][]MazeSquare, startX float32, startY float32, endX float32, endY float32) []MazeSquare {
	markUnvisited(gameGridDFS)
	var definiteNodes []MazeSquare

	size := len(gameGridDFS[0])

	for y := 0; y < size; y++ {

		for x := 0; x < size; x++ {

			if (gameGridDFS[y][x].XCoordinate == startX) && (gameGridDFS[y][x].YCoordinate == startY) || (gameGridDFS[y][x].XCoordinate == endX) && (gameGridDFS[y][x].YCoordinate == endY) {
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

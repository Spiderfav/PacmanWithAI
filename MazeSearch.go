package main

func calculateWeights(gameGridDFS *[8][8]MazeSquare, startX int, startY int, finishX int, finishY int) {

	markUnvisited(gameGridDFS)

	prevWeight := 0
	var nodePrevWeights []int

	var splitNodes []MazeSquare

	gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = 0

	// While the bottom right hand corner of the grid does not have a distance assigned and all the paths have not been taken

	for !gameGridDFS[int(finishX/20)-1][int(finishY/20)-1].Visited || len(splitNodes) != 0 {

		if gameGridDFS[int(startX/20)-1][int(startY/20)-1] != gameGridDFS[0][0] {
			prevWeight += 1
			gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = prevWeight
		}

		gameGridDFS[int(startX/20)-1][int(startY/20)-1].Visited = true

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasDown && !gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].Visited {
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1+1][int(startY/20)-1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)
		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasUp && !gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].Visited {
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1-1][int(startY/20)-1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasLeft && !gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].Visited {
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1][int(startY/20)-1-1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasRight && !gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].Visited {
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1][int(startY/20)-1+1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)

		}

		if len(splitNodes) != 0 {
			nodePopped := splitNodes[len(splitNodes)-1]
			splitNodes = splitNodes[:len(splitNodes)-1]

			prevWeight = nodePrevWeights[len(nodePrevWeights)-1]
			nodePrevWeights = nodePrevWeights[:len(nodePrevWeights)-1]

			startY = int(nodePopped.XCoordinate)
			startX = int(nodePopped.YCoordinate)
		}

	}
}

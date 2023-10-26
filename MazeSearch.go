package main

import (
	"fmt"
)

func calculateWeights(gameGridDFS *[8][8]MazeSquare, startX int, startY int, finishX int, finishY int) {

	markUnvisited(gameGridDFS)
	//For now, setting the start point to be 0,0 and the end point to be 7,7
	startX = 20
	startY = 20

	finishX = 160
	finishY = 160

	prevWeight := 0
	var nodePrevWeights []int

	var splitNodes []MazeSquare

	gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = 0

	fmt.Println(gameGridDFS[0][0])

	// While the bottom right hand corner of the grid does not have a distance assigned and all the paths have not been taken
	// Will change this to be the end node we want

	fmt.Println(" ")
	for !gameGridDFS[7][7].Visited || startX == 0 {

		if gameGridDFS[int(startX/20)-1][int(startY/20)-1] != gameGridDFS[0][0] {
			prevWeight += 1
			gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = prevWeight
		}

		fmt.Println("Here is current node before", gameGridDFS[int(startX/20)-1][int(startY/20)-1])
		gameGridDFS[int(startX/20)-1][int(startY/20)-1].Visited = true
		fmt.Println("Here is current node after", gameGridDFS[int(startX/20)-1][int(startY/20)-1])
		fmt.Println(" ")

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasDown && !gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].Visited {
			fmt.Println("Down node", gameGridDFS[int(startX/20)-1+1][int(startY/20)-1])
			fmt.Println(" ")
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1+1][int(startY/20)-1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)
		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasUp && !gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].Visited {
			fmt.Println("Up node", gameGridDFS[int(startX/20)-1-1][int(startY/20)-1])
			fmt.Println(" ")
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1-1][int(startY/20)-1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasLeft && !gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].Visited {
			fmt.Println("Left node", gameGridDFS[int(startX/20)-1][int(startY/20)-1-1])
			fmt.Println(" ")
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1][int(startY/20)-1-1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasRight && !gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].Visited {
			fmt.Println("Left node", gameGridDFS[int(startX/20)-1][int(startY/20)-1+1])
			fmt.Println(" ")
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1][int(startY/20)-1+1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)

		}

		fmt.Println("Here is the nodes in the array: ", splitNodes)
		fmt.Println("Here is the prevWeights in the array: ", nodePrevWeights)
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")

		//time.Sleep(2 * time.Second)

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

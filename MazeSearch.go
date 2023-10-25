package main

import (
	"fmt"
	"time"
)

func calculateWeights(gameGridDFS *[8][8]MazeSquare, startX int, startY int, finishX int, finishY int) {

	//For now, setting the start point to be 0,0 and the end point to be 7,7
	startX = 20
	startY = 20

	finishX = 160
	finishY = 160

	var splitNodes []MazeSquare

	gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = 0

	fmt.Println(gameGridDFS[0][0])

	// While the bottom right hand corner of the grid does not have a distance assigned and all the paths have not been taken
	// Will change this to be the end node we want

	fmt.Println(" ")
	for (gameGridDFS[7][7].Weight == 0) || (len(splitNodes) != 0) {
		fmt.Println(" ")
		fmt.Println(gameGridDFS[int(startX/20)-1][int(startY/20)-1])
		fmt.Println(" ")
		startXTemp := 0
		startYTemp := 0

		nodeSplitCounter := 0

		fmt.Println("Current X Start : ", startX)
		fmt.Println("Current Y Start : ", startY)
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("Current Weight : ", gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight)
		fmt.Println(" ")

		//time.Sleep(2 * time.Second)

		fmt.Println("Has Down Wall: ", gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasDown)

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasDown && gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].Weight == 0 {

			nodeSplitCounter += 1
			gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].Weight = gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight + 1
			startXTemp = int(gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].XCoordinate)
			startYTemp = int(gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].YCoordinate)
			fmt.Println("New Down Weight : ", gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].Weight)
		}

		fmt.Println("Has Up Wall: ", gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasUp)

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasUp && gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].Weight == 0 {

			nodeSplitCounter += 1
			gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].Weight = gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight + 1
			startXTemp = int(gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].XCoordinate)
			startYTemp = int(gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].YCoordinate)
			fmt.Println("New Up Weight : ", gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].Weight)
		}

		fmt.Println("Has Left Wall: ", gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasLeft)

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasLeft && gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].Weight == 0 {

			nodeSplitCounter += 1
			gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].Weight = gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight + 1
			startXTemp = int(gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].XCoordinate)
			startYTemp = int(gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].YCoordinate)
			fmt.Println("New Left Weight : ", gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].Weight)
		}

		fmt.Println("Has Right Wall: ", gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasRight)

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasRight && gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].Weight == 0 {

			nodeSplitCounter += 1
			gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].Weight = gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight + 1
			startXTemp = int(gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].XCoordinate)
			startYTemp = int(gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].YCoordinate)
			fmt.Println("New Right Weight : ", gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].Weight)
		}

		fmt.Println(" ")
		fmt.Println(gameGridDFS[int(startX/20)-1][int(startY/20)-1])
		fmt.Println(" ")

		startY = startXTemp
		startX = startYTemp

		fmt.Println("Current X End : ", startX)
		fmt.Println("Current Y End : ", startY)
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")

		time.Sleep(2 * time.Second)

		if nodeSplitCounter > 0 {
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1][int(startY/20)-1])
		}

		if len(splitNodes) != 0 && nodeSplitCounter == 0 {
			nodePopped := splitNodes[len(splitNodes)-1]
			splitNodes = splitNodes[:len(splitNodes)-1]
			startX = int(nodePopped.XCoordinate)
			startY = int(nodePopped.YCoordinate)
		}

		fmt.Println("Is weight? ", gameGridDFS[7][7].Weight == 0)
		fmt.Println("Is length? ", len(splitNodes) == 0)

	}
}

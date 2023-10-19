package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var gameGrid = CreateMaze()

func DFS() [8][8]MazeSquare {

	gridSize := 7

	// printGrid(gameGrid)

	var stack []*MazeSquare

	var nextNodeNoGrid *MazeSquare

	// Randomly selecting a node
	startPointX := rand.Intn(gridSize)
	startPointY := rand.Intn(gridSize)

	startNode := &gameGrid[startPointX][startPointY]

	// Appending node to stack
	stack = append(stack, startNode)

	for len(stack) != 0 {
		currentAllNodes := 0

		//fmt.Println("Start Node After Assignment ", startNode)

		// Marking node as visited
		gameGrid[int(startNode.YCoordinate/20)-1][int(startNode.XCoordinate/20)-1].Visited = true

		//printGrid(gameGrid)

		//fmt.Println("Start node = ", startNode)

		// Choose random direction to go in
		nextNodeNoGrid = chooseDirection(int(startNode.XCoordinate), int(startNode.YCoordinate))

		nextNode := gameGrid[int(nextNodeNoGrid.YCoordinate/20)-1][int(nextNodeNoGrid.XCoordinate/20)-1]

		//fmt.Println("Node chosen: ", nextNode)

		//fmt.Println("Is visited: ", nextNode.Visited)

		//time.Sleep(2 * time.Second)

		if nextNode.Visited {

			//fmt.Println("Oh no it's been visited!")

			currentAllNodes = 1
			startNode = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

		}

		if currentAllNodes == 1 {
			continue
		}

		//fmt.Println("Start Node Before ", startNode)

		startNode = &MazeSquare{}

		startNode = &nextNode

		//fmt.Println("Start Node After ", startNode)

		// Appending node to stack
		stack = append(stack, startNode)

		//moveOptions := [4]MazeSquare{*startNode.Left, *startNode.Down, *startNode.Right, *startNode.Up}

		// If direction picked does not have a node, pick another direction
	}

	//fmt.Println("Finished function!!!")
	return gameGrid

}

func DrawSquare(screen *ebiten.Image, squareToDraw MazeSquare) {

	fmt.Println("Drawing this square now: ", squareToDraw)

	var strokeWidth float32 = 1

	if squareToDraw.HasDown {
		vector.StrokeLine(screen, squareToDraw.XCoordinate, squareToDraw.YCoordinate+20, squareToDraw.XCoordinate+20, squareToDraw.YCoordinate+20, strokeWidth, color.Black, false)
	}

	if squareToDraw.HasRight {
		vector.StrokeLine(screen, squareToDraw.XCoordinate+20, squareToDraw.YCoordinate, squareToDraw.XCoordinate+20, squareToDraw.YCoordinate+20, strokeWidth, color.Black, false)
	}

	if squareToDraw.HasLeft {
		vector.StrokeLine(screen, squareToDraw.XCoordinate, squareToDraw.YCoordinate, squareToDraw.XCoordinate, squareToDraw.YCoordinate+20, strokeWidth, color.Black, false)
	}

	if squareToDraw.HasUp {
		vector.StrokeLine(screen, squareToDraw.XCoordinate, squareToDraw.YCoordinate, squareToDraw.XCoordinate+20, squareToDraw.YCoordinate, strokeWidth, color.Black, false)
	}

	fmt.Println("Finished the function!")
	fmt.Println(" ")

}

func chooseDirection(x int, y int) *MazeSquare {
	startNode := gameGrid[(y/20)-1][(x/20)-1]

	//fmt.Println("Getting from Game grid = ", startNode)

	var options []int

	var direction *MazeSquare

	directionNumber := 0

	if ((y / 20) - 1) != 7 {
		if (gameGrid[((y/20)-1)+1][(x/20)-1] != MazeSquare{}) && !gameGrid[((y/20)-1)+1][(x/20)-1].Visited {
			//fmt.Println("Is Visited in Direction Down: ", startNode.Down.Visited)
			options = append(options, 1)
		}
	}

	if ((y / 20) - 1) != 0 {
		if (gameGrid[((y/20)-1)-1][(x/20)-1] != MazeSquare{}) && !gameGrid[((y/20)-1)-1][(x/20)-1].Visited {
			//fmt.Println("Is Visited in Direction Up: ", startNode.Up.Visited)

			options = append(options, 3)
		}
	}

	if ((x / 20) - 1) != 7 {
		if (gameGrid[(y/20)-1][((x/20)-1)+1] != MazeSquare{}) && !gameGrid[(y/20)-1][((x/20)-1)+1].Visited {
			//fmt.Println("Is Visited in Direction Right: ", startNode.Right.Visited)

			options = append(options, 2)
		}
	}

	if ((x / 20) - 1) != 0 {
		if (gameGrid[(y/20)-1][((x/20)-1)-1] != MazeSquare{}) && !gameGrid[(y/20)-1][((x/20)-1)-1].Visited {
			//fmt.Println("Is Visited in Direction Left: ", startNode.Left.Visited)

			options = append(options, 0)
		}
	}

	//fmt.Println("Array: ", options)

	if len(options) == 0 {
		return &gameGrid[(y/20)-1][(x/20)-1]
	}

	nodeChosenPos := rand.Intn(len(options))

	directionNumber = options[nodeChosenPos]

	switch directionNumber {

	case 0:
		// if startNode.Left == nil {
		// 	direction = MazeSquare{}
		// 	break
		// }
		//fmt.Println("Left: ", startNode.Left)

		direction = startNode.Left

		//fmt.Println("Direcion: ", direction)

		gameGrid[(y/20)-1][(x/20)-1].HasLeft = false
		gameGrid[(y/20)-1][((x/20)-1)-1].HasRight = false

	case 1:
		// if startNode.Down == nil {
		// 	direction = MazeSquare{}
		// 	break
		// }

		//fmt.Println("Down: ", startNode.Down)

		direction = startNode.Down
		//fmt.Println("Direcion: ", direction)
		gameGrid[(y/20)-1][(x/20)-1].HasDown = false
		gameGrid[((y/20)-1)+1][(x/20)-1].HasUp = false

	case 2:
		// if startNode.Right == nil {
		// 	direction = MazeSquare{}
		// 	break
		// }
		//fmt.Println("Right: ", startNode.Right)
		direction = startNode.Right
		//fmt.Println("Direcion: ", direction)
		gameGrid[(y/20)-1][(x/20)-1].HasRight = false
		gameGrid[(y/20)-1][((x/20)-1)+1].HasLeft = false

	case 3:

		// if startNode.Up == nil {
		// 	direction = MazeSquare{}
		// 	break
		// }
		//fmt.Println("Up: ", startNode.Up)
		direction = startNode.Up
		//fmt.Println("Direcion: ", direction)
		gameGrid[(y/20)-1][(x/20)-1].HasUp = false
		gameGrid[((y/20)-1)-1][(x/20)-1].HasDown = false

	}

	// }

	return direction
}

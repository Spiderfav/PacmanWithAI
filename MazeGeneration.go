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

		// Marking node as visited
		gameGrid[int(startNode.YCoordinate/20)-1][int(startNode.XCoordinate/20)-1].Visited = true

		// Choose random direction to go in
		nextNodeNoGrid = chooseDirection(int(startNode.XCoordinate), int(startNode.YCoordinate))

		nextNode := gameGrid[int(nextNodeNoGrid.YCoordinate/20)-1][int(nextNodeNoGrid.XCoordinate/20)-1]

		if nextNode.Visited {
			currentAllNodes = 1
			startNode = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

		}

		if currentAllNodes == 1 {
			continue
		}

		startNode = &MazeSquare{}

		startNode = &nextNode

		// Appending node to stack
		stack = append(stack, startNode)
	}

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

	var options []int

	var direction *MazeSquare

	directionNumber := 0

	if ((y / 20) - 1) != 7 {
		if (gameGrid[((y/20)-1)+1][(x/20)-1] != MazeSquare{}) && !gameGrid[((y/20)-1)+1][(x/20)-1].Visited {
			options = append(options, 1)
		}
	}

	if ((y / 20) - 1) != 0 {
		if (gameGrid[((y/20)-1)-1][(x/20)-1] != MazeSquare{}) && !gameGrid[((y/20)-1)-1][(x/20)-1].Visited {

			options = append(options, 3)
		}
	}

	if ((x / 20) - 1) != 7 {
		if (gameGrid[(y/20)-1][((x/20)-1)+1] != MazeSquare{}) && !gameGrid[(y/20)-1][((x/20)-1)+1].Visited {

			options = append(options, 2)
		}
	}

	if ((x / 20) - 1) != 0 {
		if (gameGrid[(y/20)-1][((x/20)-1)-1] != MazeSquare{}) && !gameGrid[(y/20)-1][((x/20)-1)-1].Visited {

			options = append(options, 0)
		}
	}

	if len(options) == 0 {
		return &gameGrid[(y/20)-1][(x/20)-1]
	}

	nodeChosenPos := rand.Intn(len(options))

	directionNumber = options[nodeChosenPos]

	switch directionNumber {

	case 0:
		direction = startNode.Left

		gameGrid[(y/20)-1][(x/20)-1].HasLeft = false
		gameGrid[(y/20)-1][((x/20)-1)-1].HasRight = false

	case 1:

		direction = startNode.Down

		gameGrid[(y/20)-1][(x/20)-1].HasDown = false
		gameGrid[((y/20)-1)+1][(x/20)-1].HasUp = false

	case 2:

		direction = startNode.Right

		gameGrid[(y/20)-1][(x/20)-1].HasRight = false
		gameGrid[(y/20)-1][((x/20)-1)+1].HasLeft = false

	case 3:

		direction = startNode.Up

		gameGrid[(y/20)-1][(x/20)-1].HasUp = false
		gameGrid[((y/20)-1)-1][(x/20)-1].HasDown = false

	}

	return direction
}

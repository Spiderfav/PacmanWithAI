package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var gameGrid = CreateGrid()

func DFS() [8][8]MazeSquare {

	gridSize := 7

	var stack []*MazeSquare

	var nextNodeNoGrid *MazeSquare

	// Randomly selecting a node from the grid
	startPointX := rand.Intn(gridSize)
	startPointY := rand.Intn(gridSize)

	// Selected node chosen
	startNode := &gameGrid[startPointX][startPointY]

	// Appending node to stack
	stack = append(stack, startNode)

	// While the stack of nodes is not empty; While we have not visited every node
	for len(stack) != 0 {
		currentAllNodes := 0

		// Marking node as visited
		gameGrid[int(startNode.YCoordinate/20)-1][int(startNode.XCoordinate/20)-1].Visited = true

		// Choose random direction to go in
		nextNodeNoGrid = chooseDirection(int(startNode.XCoordinate), int(startNode.YCoordinate))

		// Get the node for the direction we want to go in
		nextNode := gameGrid[int(nextNodeNoGrid.YCoordinate/20)-1][int(nextNodeNoGrid.XCoordinate/20)-1]

		// If the node we picked has already been visited, pop off the stack until one that hasn't been visited is chosen
		if nextNode.Visited {
			currentAllNodes = 1
			startNode = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

		}

		// Resets the while loop to get a new node
		// If all the nodes have been popped off the stack, exits the loop
		if currentAllNodes == 1 {
			continue
		}

		// Clearing the variable
		startNode = &MazeSquare{}

		// Assigning the new startNode to be the direction chosen
		startNode = &nextNode

		// Appending node to stack
		stack = append(stack, startNode)
	}

	return gameGrid

}

// This function draws a given square to the screen
func DrawSquare(screen *ebiten.Image, squareToDraw MazeSquare) {
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

}

// This function, given an X and Y co-ordinate from a MazeNode, choses a random direction to go in
func chooseDirection(x int, y int) *MazeSquare {
	startNode := gameGrid[(y/20)-1][(x/20)-1]

	var options []int

	var direction *MazeSquare

	directionNumber := 0

	// These if blocks check if the MazeSquare chosen is not an edge and that its neighbours are not empty or visited

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

	// If there are no options to choose
	if len(options) == 0 {
		// Return the same object
		return &gameGrid[(y/20)-1][(x/20)-1]
	}

	// Choose a random direction out of the available ones to go to
	nodeChosenPos := rand.Intn(len(options))

	directionNumber = options[nodeChosenPos]

	// Given the direction chosen, break the walls between the two nodes

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

func markUnvisited(gameGridDFS *[8][8]MazeSquare) {

	for y := 0; y < 8; y++ {

		for x := 0; x < 8; x++ {

			gameGridDFS[y][x].Visited = false

		}

	}

}

package main

import (
	"math/rand"
)

// This function uses randomized DFS to generate a maze
func DFS(size int) [][]MazeSquare {

	var gameGrid = CreateGrid(size)

	size = size - 1

	gridSize := size

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
		nextNodeNoGrid = chooseDirection(int(startNode.XCoordinate), int(startNode.YCoordinate), size, gameGrid)

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

// This function, given an X and Y co-ordinate from a MazeNode, choses a random direction to go in
func chooseDirection(x int, y int, size int, gameGrid [][]MazeSquare) *MazeSquare {
	startNode := gameGrid[(y/20)-1][(x/20)-1]

	var options []int

	var direction *MazeSquare

	directionNumber := 0

	// These if blocks check if the MazeSquare chosen is not an edge and that its neighbours are not empty or visited

	if ((y / 20) - 1) != size {
		if (gameGrid[((y/20)-1)+1][(x/20)-1] != MazeSquare{}) && !gameGrid[((y/20)-1)+1][(x/20)-1].Visited {
			options = append(options, 1)
		}
	}

	if ((y / 20) - 1) != 0 {
		if (gameGrid[((y/20)-1)-1][(x/20)-1] != MazeSquare{}) && !gameGrid[((y/20)-1)-1][(x/20)-1].Visited {

			options = append(options, 3)
		}
	}

	if ((x / 20) - 1) != size {
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

// This simple function is run before any pathfinding algorithm to make sure that the nodes are marked unvisited
func markUnvisited(gameGridDFS [][]MazeSquare, size int) {

	for y := 0; y < size; y++ {

		for x := 0; x < size; x++ {

			gameGridDFS[y][x].Visited = false
			gameGridDFS[y][x].Weight = 0

		}

	}

}

// This functions adds weights to a specific square in the grid
func addWeights(gameGridDFS [][]MazeSquare, obstacle int) {
	//xValue := rand.Intn(len(gameGridDFS[0]))
	//yvalue := rand.Intn(len(gameGridDFS[0]))

	gameGridDFS[0][1].Weight = obstacle
}

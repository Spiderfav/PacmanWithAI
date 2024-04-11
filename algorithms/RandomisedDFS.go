package algorithms

import (
	"math/rand"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// This function uses randomized DFS to generate a maze
func DFS(size int, oldDFS [][]mazegrid.MazeSquare, squareSize int) [][]mazegrid.MazeSquare {

	var gameGrid [][]mazegrid.MazeSquare

	if oldDFS == nil {
		gameGrid = mazegrid.CreateGrid(size, squareSize)

	} else {
		gameGrid = oldDFS
	}

	size = size - 1

	gridSize := size

	var stack []*mazegrid.MazeSquare

	var nextNodeNoGrid *mazegrid.MazeSquare

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
		startNode.Visited = true

		// Choose random direction to go in
		nextNodeNoGrid = chooseDirection(int(startNode.NodePosition.XCoordinate), int(startNode.NodePosition.YCoordinate), size, gameGrid, squareSize)

		// Get the node for the direction we want to go in
		nextNode := &gameGrid[int(nextNodeNoGrid.NodePosition.YCoordinate/float32(squareSize))-1][int(nextNodeNoGrid.NodePosition.XCoordinate/float32(squareSize))-1]

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
		startNode = &mazegrid.MazeSquare{}

		// Assigning the new startNode to be the direction chosen
		startNode = nextNode

		// Appending node to stack
		stack = append(stack, startNode)
	}

	// Adding random weights in the gamegrid (They are considered pellots)
	AddWeights(gameGrid)

	return gameGrid

}

// This function, given an X and Y co-ordinate from a MazeNode, choses a random direction to go in
func chooseDirection(x int, y int, size int, gameGrid [][]mazegrid.MazeSquare, squareSize int) *mazegrid.MazeSquare {
	startNode := &gameGrid[(y/squareSize)-1][(x/squareSize)-1]

	var options []int

	var direction *mazegrid.MazeSquare

	directionNumber := 0

	// These if blocks check if the MazeSquare chosen is not an edge and that its neighbours are not empty or visited

	// If the Y value is not in the furthest edge
	if ((y / squareSize) - 1) != size {
		node := &gameGrid[((y/squareSize)-1)+1][(x/squareSize)-1]
		if (*node != mazegrid.MazeSquare{}) && !node.Visited {
			options = append(options, 1)
		}
	}

	// If the Y value is not in the closest edge
	if ((y / squareSize) - 1) != 0 {
		node := &gameGrid[((y/squareSize)-1)-1][(x/squareSize)-1]

		if (*node != mazegrid.MazeSquare{}) && !node.Visited {

			options = append(options, 3)
		}
	}

	// If the X value is not in the furthest edge
	if ((x / squareSize) - 1) != size {
		node := &gameGrid[(y/squareSize)-1][((x/squareSize)-1)+1]

		if (*node != mazegrid.MazeSquare{}) && !node.Visited {

			options = append(options, 2)
		}
	}

	// If the X value is not in the closest edge
	if ((x / squareSize) - 1) != 0 {
		node := &gameGrid[(y/squareSize)-1][((x/squareSize)-1)-1]

		if (*node != mazegrid.MazeSquare{}) && !node.Visited {

			options = append(options, 0)
		}
	}

	// If there are no options to choose
	if len(options) == 0 {
		// Return the same object
		return &gameGrid[(y/squareSize)-1][(x/squareSize)-1]
	}

	// Choose a random direction out of the available ones to go to
	nodeChosenPos := rand.Intn(len(options))

	directionNumber = options[nodeChosenPos]

	// Given the direction chosen, break the walls between the two nodes

	switch directionNumber {

	// Going left
	case 0:
		direction = mazegrid.PosToNode(gameGrid, startNode.Walls.Left, squareSize)

		// Break the left wall of the current node
		startNode.HasWalls.HasLeft = false

		// Break the right wall of the chosen node
		direction.HasWalls.HasRight = false

	// Going down
	case 1:

		direction = mazegrid.PosToNode(gameGrid, startNode.Walls.Down, squareSize)

		// Break the down wall of the current node
		startNode.HasWalls.HasDown = false

		// Break the up wall of the chosen node
		direction.HasWalls.HasUp = false

	// Going right
	case 2:

		direction = mazegrid.PosToNode(gameGrid, startNode.Walls.Right, squareSize)

		// Break the right wall of the current node
		startNode.HasWalls.HasRight = false

		// Break the left wall of the chosen node
		direction.HasWalls.HasLeft = false

	// Goinf up
	case 3:

		direction = mazegrid.PosToNode(gameGrid, startNode.Walls.Up, squareSize)

		// Break the up wall of the current node
		startNode.HasWalls.HasUp = false

		// Break the down wall of the chosen node
		direction.HasWalls.HasDown = false

	}

	// Return the direction chosen
	return direction
}

// Generates a maze given the size of the maze and the square size
func CreateMaze(mazeSize int, squareSize int) mazegrid.Maze {
	// Creating the maze by aplying DFS twice
	oldGameGridDFS := DFS(mazeSize, nil, squareSize)
	MarkUnvisited(oldGameGridDFS, false)
	gameGridDFS := DFS(mazeSize, oldGameGridDFS, squareSize)
	return mazegrid.Maze{Size: mazeSize, Grid: gameGridDFS, Pellots: mazegrid.GetPellotsPos(gameGridDFS)}

}

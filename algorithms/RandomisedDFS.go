package algorithms

import (
	"fmt"
	"math/rand"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// This function uses randomized DFS to generate a maze
func DFS(size int, oldDFS [][]mazegrid.MazeSquare) [][]mazegrid.MazeSquare {

	var gameGrid [][]mazegrid.MazeSquare

	if oldDFS == nil {
		gameGrid = mazegrid.CreateGrid(size)

	} else {
		gameGrid = oldDFS
	}

	DebugMaze(gameGrid)

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
		gameGrid[int(startNode.NodePosition.YCoordinate/20)-1][int(startNode.NodePosition.XCoordinate/20)-1].Visited = true

		// Choose random direction to go in
		nextNodeNoGrid = chooseDirection(int(startNode.NodePosition.XCoordinate), int(startNode.NodePosition.YCoordinate), size, gameGrid)

		// Get the node for the direction we want to go in
		nextNode := gameGrid[int(nextNodeNoGrid.NodePosition.YCoordinate/20)-1][int(nextNodeNoGrid.NodePosition.XCoordinate/20)-1]

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
		startNode = &nextNode

		// Appending node to stack
		stack = append(stack, startNode)
	}

	return gameGrid

}

// This function, given an X and Y co-ordinate from a MazeNode, choses a random direction to go in
func chooseDirection(x int, y int, size int, gameGrid [][]mazegrid.MazeSquare) *mazegrid.MazeSquare {
	startNode := gameGrid[(y/20)-1][(x/20)-1]

	var options []int

	var direction *mazegrid.MazeSquare

	directionNumber := 0

	// These if blocks check if the MazeSquare chosen is not an edge and that its neighbours are not empty or visited

	if ((y / 20) - 1) != size {
		if (gameGrid[((y/20)-1)+1][(x/20)-1] != mazegrid.MazeSquare{}) && !gameGrid[((y/20)-1)+1][(x/20)-1].Visited {
			options = append(options, 1)
		}
	}

	if ((y / 20) - 1) != 0 {
		if (gameGrid[((y/20)-1)-1][(x/20)-1] != mazegrid.MazeSquare{}) && !gameGrid[((y/20)-1)-1][(x/20)-1].Visited {

			options = append(options, 3)
		}
	}

	if ((x / 20) - 1) != size {
		if (gameGrid[(y/20)-1][((x/20)-1)+1] != mazegrid.MazeSquare{}) && !gameGrid[(y/20)-1][((x/20)-1)+1].Visited {

			options = append(options, 2)
		}
	}

	if ((x / 20) - 1) != 0 {
		if (gameGrid[(y/20)-1][((x/20)-1)-1] != mazegrid.MazeSquare{}) && !gameGrid[(y/20)-1][((x/20)-1)-1].Visited {

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
		direction = mazegrid.PosToNode(gameGrid, startNode.Walls.Left)

		//time.Sleep(5 * time.Second)

		//direction = &gameGrid[int(startNode.Walls.Left.XCoordinate/20)-1][int(startNode.Walls.Left.YCoordinate/20)-1]

		gameGrid[(y/20)-1][(x/20)-1].HasWalls.HasLeft = false
		gameGrid[(y/20)-1][((x/20)-1)-1].HasWalls.HasRight = false

	case 1:

		direction = mazegrid.PosToNode(gameGrid, startNode.Walls.Down)
		//direction = &gameGrid[int(startNode.Walls.Down.XCoordinate/20)-1][int(startNode.Walls.Down.YCoordinate/20)-1]

		gameGrid[(y/20)-1][(x/20)-1].HasWalls.HasDown = false
		gameGrid[((y/20)-1)+1][(x/20)-1].HasWalls.HasUp = false

	case 2:

		direction = mazegrid.PosToNode(gameGrid, startNode.Walls.Right)
		//direction = &gameGrid[int(startNode.Walls.Right.XCoordinate/20)-1][int(startNode.Walls.Right.YCoordinate/20)-1]

		gameGrid[(y/20)-1][(x/20)-1].HasWalls.HasRight = false
		gameGrid[(y/20)-1][((x/20)-1)+1].HasWalls.HasLeft = false

	case 3:

		direction = mazegrid.PosToNode(gameGrid, startNode.Walls.Up)
		//direction = &gameGrid[int(startNode.Walls.Up.XCoordinate/20)-1][int(startNode.Walls.Up.YCoordinate/20)-1]

		gameGrid[(y/20)-1][(x/20)-1].HasWalls.HasUp = false
		gameGrid[((y/20)-1)-1][(x/20)-1].HasWalls.HasDown = false

	}

	return direction
}

func DebugMaze(m [][]mazegrid.MazeSquare) {
	for j := 0; j < len(m[0]); j++ {

		for i := 0; i < len(m[0]); i++ {
			fmt.Println("Node: ", m[j][i])

		}

		fmt.Println(" ")
	}
}

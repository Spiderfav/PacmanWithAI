package mazegrid

// This functions creates and returns an empty MazeSquare object
func CreateBlankSquare(squareSize int) MazeSquare {
	return MazeSquare{Position{float32(squareSize), float32(squareSize)}, Direction{}, HasDirection{true, true, true, true}, false, 0, 0, true, false}

}

// Takes 2 parameters: The Game grid, position
// This function, given the grid and a position, will return the MazeSquare from the game grid of the given position
func PosToNode(gameGrid [][]MazeSquare, p Position, squareSize int) *MazeSquare {
	yCoord := int(p.YCoordinate/float32(squareSize)) - 1
	xCoord := int(p.XCoordinate/float32(squareSize)) - 1

	return &gameGrid[yCoord][xCoord]
}

// This function creates a grid of squareSize*squareSize MazeSquares, each with pointers to its direct neighbours
func CreateGrid(size int, squareSize int) [][]MazeSquare {

	// Define the size of the grid

	gameGrid := make([][]MazeSquare, size)
	for i := range gameGrid {
		gameGrid[i] = make([]MazeSquare, size)
	}

	var x, y float32

	var squareLengthX, squareLengthY float32 = float32(squareSize), float32(squareSize)

	// Even though we are using X and Y co-ordinates for the objects, in an array sense, it will be Y and X

	for y = 0; y < float32(size); y++ {

		for x = 0; x < float32(size); x++ {

			// + 1 is used to get an actual X and Y value for the starting point as X could be 0
			positionOfNode := Position{squareLengthX * (x + 1), squareLengthY * (y + 1)}

			nodeWalls := Direction{}

			nodeHasWalls := HasDirection{true, true, true, true}

			// Using i + 1 and j + 1 as this is calculating the square size and as it starts by 0, we need to add one to the normal counter
			var square = MazeSquare{positionOfNode, nodeWalls, nodeHasWalls, false, 0, 0, true, false}

			// Setting the game grid values to the MazeSquare object
			gameGrid[int(y)][int(x)] = square

			// If x>0, then we can assign left and right neighbours
			if x > 0 {
				gameGrid[int(y)][int(x)].Walls.Left = gameGrid[int(y)][int(x-1)].NodePosition
				gameGrid[int(y)][int(x-1)].Walls.Right = gameGrid[int(y)][int(x)].NodePosition
			}

			// If y>0, then we can assign up and down neighbours
			if y > 0 {
				gameGrid[int(y)][int(x)].Walls.Up = gameGrid[int(y-1)][int(x)].NodePosition
				gameGrid[int(y-1)][int(x)].Walls.Down = gameGrid[int(y)][int(x)].NodePosition

			}

		}

	}

	return gameGrid
}

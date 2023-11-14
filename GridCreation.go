package main

// The Mazequare Object
type MazeSquare struct {
	XCoordinate float32
	YCoordinate float32
	Left        *MazeSquare
	HasLeft     bool
	Down        *MazeSquare
	HasDown     bool
	Right       *MazeSquare
	HasRight    bool
	Up          *MazeSquare
	HasUp       bool
	Visited     bool
	Weight      int
}

// This function creates a grid of 8*8 MazeSquares, each with pointers to its direct neighbours
func CreateGrid(size int) [][]MazeSquare {

	// Define the size of the grid

	gameGrid := make([][]MazeSquare, size)
	for i := range gameGrid {
		gameGrid[i] = make([]MazeSquare, size)
	}

	var x, y float32

	var squareLengthX, squareLengthY float32 = 20, 20

	// Even though we are using X and Y co-ordinates for the objects, in an array sense, it will be Y and X

	for y = 0; y < float32(size); y++ {

		for x = 0; x < float32(size); x++ {

			// Using i + 1 and j + 1 as this is calculating the square size and as it starts by 0, we need to add one to the normal counter
			var square = MazeSquare{squareLengthX * (x + 1), squareLengthY * (y + 1), nil, true, nil, true, nil, true, nil, true, false, 0}

			// Setting the game grid values to the MazeSquare object
			gameGrid[int(y)][int(x)] = square

			// If x>0, then we can assign left and right neighbours
			if x > 0 {
				gameGrid[int(y)][int(x)].Left = &gameGrid[int(y)][int(x-1)]
				gameGrid[int(y)][int(x-1)].Right = &gameGrid[int(y)][int(x)]
			}

			// If y>0, then we can assign up and down neighbours
			if y > 0 {
				gameGrid[int(y)][int(x)].Up = &gameGrid[int(y-1)][int(x)]
				gameGrid[int(y-1)][int(x)].Down = &gameGrid[int(y)][int(x)]

			}

		}

	}

	return gameGrid
}

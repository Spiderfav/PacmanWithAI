package mazegrid

// The Position object is used for a MazeSquare object
// It contains the X and Y Coordinates of an MazeSquare object
type Position struct {
	XCoordinate float32
	YCoordinate float32
}

// The Direction object is used for a MazeSquare object
// It contains a pointer to all the neightbours of the MazeSquare object
type Direction struct {
	Left  *MazeSquare
	Down  *MazeSquare
	Right *MazeSquare
	Up    *MazeSquare
}

// The HasDirection object is used for a MazeSquare object
// It contains the walls that the MazeSquare object has
type HasDirection struct {
	HasLeft  bool
	HasDown  bool
	HasRight bool
	HasUp    bool
}

// The Mazequare Object
type MazeSquare struct {
	NodePosition  Position
	Walls         Direction
	HasWalls      HasDirection
	Visited       bool
	Weight        int
	NumberOfWalls int
}

func CreateBlankSquare() MazeSquare {
	return MazeSquare{Position{20, 20}, Direction{nil, nil, nil, nil}, HasDirection{true, true, true, true}, false, 0, 4}

}

// This function counts the walls of a node
func CountWalls(x MazeSquare) int {
	count := 0

	if x.HasWalls.HasLeft {
		count += 1
	}

	if x.HasWalls.HasRight {
		count += 1
	}

	if x.HasWalls.HasUp {
		count += 1
	}

	if x.HasWalls.HasDown {
		count += 1
	}

	return count
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
			positionOfNode := Position{squareLengthX * (x + 1), squareLengthY * (y + 1)}

			nodeWalls := Direction{nil, nil, nil, nil}

			nodeHasWalls := HasDirection{true, true, true, true}

			// Using i + 1 and j + 1 as this is calculating the square size and as it starts by 0, we need to add one to the normal counter
			var square = MazeSquare{positionOfNode, nodeWalls, nodeHasWalls, false, 0, 4}

			// Setting the game grid values to the MazeSquare object
			gameGrid[int(y)][int(x)] = square

			// If x>0, then we can assign left and right neighbours
			if x > 0 {
				gameGrid[int(y)][int(x)].Walls.Left = &gameGrid[int(y)][int(x-1)]
				gameGrid[int(y)][int(x-1)].Walls.Right = &gameGrid[int(y)][int(x)]
			}

			// If y>0, then we can assign up and down neighbours
			if y > 0 {
				gameGrid[int(y)][int(x)].Walls.Up = &gameGrid[int(y-1)][int(x)]
				gameGrid[int(y-1)][int(x)].Walls.Down = &gameGrid[int(y)][int(x)]

			}

		}

	}

	return gameGrid
}
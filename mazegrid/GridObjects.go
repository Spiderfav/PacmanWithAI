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
	Left  Position
	Down  Position
	Right Position
	Up    Position
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
	NodePosition   Position
	Walls          Direction
	HasWalls       HasDirection
	Visited        bool
	Weight         float64
	Heuristic      float64
	HasPellot      bool
	HasSuperPellot bool
}

// This function counts the walls of a node and returns the number of walls as an integer
func (x *MazeSquare) CountWalls() int {
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

// The Maze object contains the size of the maze and the grid of the maze
type Maze struct {
	Size    int
	Grid    [][]MazeSquare
	Pellots []Position
}

// This function, given a game grid, returns an array with all the position of the pellots on the map
func GetPellotsPos(gameGridDFS [][]MazeSquare) []Position {

	var pellots []Position

	size := len(gameGridDFS[0])

	for y := 0; y < size; y++ {

		for x := 0; x < size; x++ {

			if gameGridDFS[y][x].HasPellot || gameGridDFS[y][x].HasSuperPellot {
				pellots = append(pellots, gameGridDFS[y][x].NodePosition)
			}

		}

	}

	return pellots

}

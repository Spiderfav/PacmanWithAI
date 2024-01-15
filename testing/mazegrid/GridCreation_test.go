package testing

import (
	"testing"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

func TestCreateBlankSquare(t *testing.T) {
	blankSquare := mazegrid.CreateBlankSquare()

	if blankSquare.NumberOfWalls == 0 {
		t.Errorf("Result was incorrect, got start node X Coordinate: %d, want: %s.", blankSquare.NumberOfWalls, "4")

	}
}

func TestCountWalls(t *testing.T) {
	blankSquare := mazegrid.CreateBlankSquare()
	squareTest := mazegrid.MazeSquare{NodePosition: mazegrid.Position{XCoordinate: 20, YCoordinate: 20}, Walls: mazegrid.Direction{Left: &blankSquare, Down: nil, Right: nil, Up: nil}, HasWalls: mazegrid.HasDirection{HasLeft: false, HasDown: true, HasRight: true, HasUp: true}, Visited: false, Weight: 0, NumberOfWalls: 4}

	if mazegrid.CountWalls(squareTest) == 4 {
		t.Errorf("Result was incorrect, got start node X Coordinate: %d, want: %s.", blankSquare.NumberOfWalls, "3")

	}

}

func TestCreateGrid(t *testing.T) {
	grid := mazegrid.CreateGrid(8)

	if grid[7][7].NodePosition.XCoordinate != 160 {
		t.Errorf("Result was incorrect, got start node X Coordinate: %f, want: %s.", grid[7][7].NodePosition.XCoordinate, "160")

	}

}

func BenchmarkGrid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mazegrid.CreateGrid(10000)
	}
}

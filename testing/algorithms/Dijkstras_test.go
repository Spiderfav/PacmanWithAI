package testing

import (
	"testing"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

func TestDijkstras(t *testing.T) {
	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(256)
	path := algorithms.Dijkstras(gameGridDFS, 20, 20, 256, 256)

	if len(path) == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(path), ">0")
	}
}

func TestDijkstrasStartEnd(t *testing.T) {
	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(256)
	path := algorithms.Dijkstras(gameGridDFS, 20, 20, 5120, 5120)

	if path[0].NodePosition.XCoordinate != 20 || path[0].NodePosition.YCoordinate != 20 {
		t.Errorf("Result was incorrect, got: %f, want: %s.", path[0].NodePosition.XCoordinate, "20")

	} else if (path[len(path)-1].NodePosition.XCoordinate) != 5120 || (path[len(path)-1].NodePosition.YCoordinate) != 5120 {
		t.Errorf("Result was incorrect, got: %f, want: %s.", path[len(path)-1].NodePosition.XCoordinate, "5120")
	}
}

func BenchmarkDijkstras(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(mazeSizeOriginal)
		algorithms.Dijkstras(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
	}
}

package testing

import (
	"testing"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

func TestBFS(t *testing.T) {
	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(256, nil, 20)
	path := algorithms.BFS(gameGridDFS, 20, 20, 256, 256, 20)

	if len(path) == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(path), ">0")
	}
}

func TestBFSStartEnd(t *testing.T) {
	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(256, nil, 20)
	path := algorithms.BFS(gameGridDFS, 20, 20, 5120, 5120, 20)

	if path[0].NodePosition.XCoordinate != 5120 || path[0].NodePosition.YCoordinate != 5120 {
		t.Errorf("Result was incorrect, got: %f, want: %s.", path[0].NodePosition.XCoordinate, "5120")

	} else if (path[len(path)-1].NodePosition.XCoordinate) != 20 || (path[len(path)-1].NodePosition.YCoordinate) != 20 {
		t.Errorf("Result was incorrect, got: %f, want: %s.", path[len(path)-1].NodePosition.XCoordinate, "20")
	}
}

func BenchmarkBFS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(mazeSizeOriginal, nil, 20)
		algorithms.BFS(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal, 20)
	}
}

package testing

import (
	"testing"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

func TestEuclidean(t *testing.T) {
	distance := algorithms.EuclideanDistance(0, 0, 2, 0)

	if distance != 2 {
		t.Errorf("Result was incorrect, got: %f, want: %s.", distance, "2")
	}
}

func TestAStarSize(t *testing.T) {
	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(256, nil)
	path := algorithms.AStar(gameGridDFS, 20, 20, 256, 256)

	if len(path) == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(path), ">0")
	}
}

func TestAStarStartEnd(t *testing.T) {
	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(256, nil)
	path := algorithms.AStar(gameGridDFS, 20, 20, 5120, 5120)

	if path[0].NodePosition.XCoordinate != 20 || path[0].NodePosition.YCoordinate != 20 {
		t.Errorf("Result was incorrect, got: %f, want: %s.", path[0].NodePosition.XCoordinate, "20")

	} else if (path[len(path)-1].NodePosition.XCoordinate) != 5120 || (path[len(path)-1].NodePosition.YCoordinate) != 5120 {
		t.Errorf("Result was incorrect, got: %f, want: %s.", path[len(path)-1].NodePosition.XCoordinate, "5120")
	}
}

func BenchmarkEuclidean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithms.EuclideanDistance(2, 2, 4, 4)
	}
}

func BenchmarkAStar(b *testing.B) {
	var gridSize = 1028
	for i := 0; i < b.N; i++ {
		var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(gridSize, nil)
		algorithms.AStar(gameGridDFS, 20, 20, 20*gridSize, 20*gridSize)
	}
}

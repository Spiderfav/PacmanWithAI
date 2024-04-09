package testing

import (
	"testing"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

func TestCountWalls(t *testing.T) {
	node := mazegrid.CreateBlankSquare(20)

	if node.CountWalls() != 4 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", node.CountWalls(), "4")
	}
}

func TestDFS(t *testing.T) {
	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(256, nil, 20)

	if gameGridDFS[0][0].CountWalls() == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", gameGridDFS[0][0].CountWalls(), ">= 2")
	}
}

func BenchmarkDFS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithms.DFS(1028, nil, 20)
	}
}

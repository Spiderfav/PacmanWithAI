package testing

import (
	"testing"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

func TestCountWalls(t *testing.T) {
	node := mazegrid.CreateBlankSquare()

	if mazegrid.CountWalls(node) != 4 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", mazegrid.CountWalls(node), "4")
	}
}

func TestDFS(t *testing.T) {
	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(256, nil)

	if mazegrid.CountWalls(gameGridDFS[0][0]) == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", mazegrid.CountWalls(gameGridDFS[0][0]), ">= 2")
	}
}

func BenchmarkDFS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithms.DFS(1028, nil)
	}
}

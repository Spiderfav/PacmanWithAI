package algorithms

import (
	"testing"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

var mazeSizeOriginal = 4

func TestOneNode(t *testing.T) {
	var pathToTest = []mazegrid.MazeSquare{mazegrid.CreateBlankSquare()}

	path, weight := AbsolutePath(pathToTest)
	if len(path) != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(path), "1")
	} else if weight != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", weight, "0")
	}
}

func TestSmallPath(t *testing.T) {
	var gameGridDFS [][]mazegrid.MazeSquare = DFS(mazeSizeOriginal)

	var dijkstrasPath = Dijkstras(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
	var absolutePathDijkstras, weightDijkstras = AbsolutePath(dijkstrasPath)

	if len(absolutePathDijkstras) == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(absolutePathDijkstras), "Path Length = 0")
	} else if weightDijkstras == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", weightDijkstras, "Number = 0 ")
	} else if len(absolutePathDijkstras) == mazeSizeOriginal {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(absolutePathDijkstras), "Lenght of Maze ")
	} else if weightDijkstras == mazeSizeOriginal {
		t.Errorf("Result was incorrect, got: %d, want: %s.", weightDijkstras, "Length of Maze ")
	}
}

func TestMediumPath(t *testing.T) {
	mazeSizeOriginal = mazeSizeOriginal * 2
	var gameGridDFS [][]mazegrid.MazeSquare = DFS(mazeSizeOriginal)

	var dijkstrasPath = Dijkstras(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
	var absolutePathDijkstras, weightDijkstras = AbsolutePath(dijkstrasPath)

	if len(absolutePathDijkstras) == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(absolutePathDijkstras), "Path Length = 0")
	} else if weightDijkstras == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", weightDijkstras, "Number = 0 ")
	} else if len(absolutePathDijkstras) == mazeSizeOriginal {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(absolutePathDijkstras), "Lenght of Maze ")
	} else if weightDijkstras == mazeSizeOriginal {
		t.Errorf("Result was incorrect, got: %d, want: %s.", weightDijkstras, "Length of Maze ")
	}
}

func TestLargePath(t *testing.T) {
	mazeSizeOriginal = mazeSizeOriginal * mazeSizeOriginal

	var gameGridDFS [][]mazegrid.MazeSquare = DFS(mazeSizeOriginal)

	var dijkstrasPath = Dijkstras(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
	var absolutePathDijkstras, weightDijkstras = AbsolutePath(dijkstrasPath)

	if len(absolutePathDijkstras) == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(absolutePathDijkstras), "Path Length = 0")
	} else if weightDijkstras == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", weightDijkstras, "Number = 0 ")
	} else if len(absolutePathDijkstras) == mazeSizeOriginal {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(absolutePathDijkstras), "Lenght of Maze ")
	} else if weightDijkstras == mazeSizeOriginal {
		t.Errorf("Result was incorrect, got: %d, want: %s.", weightDijkstras, "Length of Maze ")
	}
}

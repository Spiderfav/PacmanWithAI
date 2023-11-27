package testing

import (
	"testing"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

var mazeSizeOriginal = 4

func TestOneNode(t *testing.T) {
	var pathToTest = []mazegrid.MazeSquare{mazegrid.CreateBlankSquare()}

	path, weight := algorithms.AbsolutePath(pathToTest)
	if len(path) != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(path), "1")
	} else if weight != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", weight, "0")
	}
}

func TestStartAndEndDikstras(t *testing.T) {
	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(mazeSizeOriginal)

	var dijkstrasPath = algorithms.Dijkstras(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
	var absolutePathDijkstras, _ = algorithms.AbsolutePath(dijkstrasPath)

	if absolutePathDijkstras[0].NodePosition.XCoordinate != float32(20*mazeSizeOriginal) && absolutePathDijkstras[0].NodePosition.YCoordinate != float32(20*mazeSizeOriginal) {
		t.Errorf("Result was incorrect, got X value: %f, Y value: %f ; Want X: %d, Y: %d.", absolutePathDijkstras[0].NodePosition.XCoordinate, absolutePathDijkstras[0].NodePosition.YCoordinate, 20*mazeSizeOriginal, 20*mazeSizeOriginal)

	} else if absolutePathDijkstras[len(absolutePathDijkstras)-1].NodePosition.XCoordinate != float32(20) && absolutePathDijkstras[len(absolutePathDijkstras)-1].NodePosition.YCoordinate != float32(20) {
		t.Errorf("Result was incorrect, got X value: %f, Y value: %f ; Want X: %d, Y: %d.", absolutePathDijkstras[0].NodePosition.XCoordinate, absolutePathDijkstras[0].NodePosition.YCoordinate, 20*mazeSizeOriginal, 20*mazeSizeOriginal)

	}
}

func TestStartAndEndAStar(t *testing.T) {
	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(mazeSizeOriginal)

	var aStarPath = algorithms.AStar(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
	var absolutePathAStar, _ = algorithms.AbsolutePath(aStarPath)

	if absolutePathAStar[0].NodePosition.XCoordinate != float32(20*mazeSizeOriginal) && absolutePathAStar[0].NodePosition.YCoordinate != float32(20*mazeSizeOriginal) {
		t.Errorf("Result was incorrect, got X value: %f, Y value: %f ; Want X: %d, Y: %d.", absolutePathAStar[0].NodePosition.XCoordinate, absolutePathAStar[0].NodePosition.YCoordinate, 20*mazeSizeOriginal, 20*mazeSizeOriginal)

	} else if absolutePathAStar[len(absolutePathAStar)-1].NodePosition.XCoordinate != float32(20) && absolutePathAStar[len(absolutePathAStar)-1].NodePosition.YCoordinate != float32(20) {
		t.Errorf("Result was incorrect, got X value: %f, Y value: %f ; Want X: %d, Y: %d.", absolutePathAStar[0].NodePosition.XCoordinate, absolutePathAStar[0].NodePosition.YCoordinate, 20*mazeSizeOriginal, 20*mazeSizeOriginal)

	}
}

func TestSmallPathDijkstras(t *testing.T) {
	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(mazeSizeOriginal)

	var dijkstrasPath = algorithms.Dijkstras(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
	var absolutePathDijkstras, weightDijkstras = algorithms.AbsolutePath(dijkstrasPath)

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

func TestMediumPathDijkstras(t *testing.T) {
	mazeSizeOriginal = mazeSizeOriginal * 2
	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(mazeSizeOriginal)

	var dijkstrasPath = algorithms.Dijkstras(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
	var absolutePathDijkstras, weightDijkstras = algorithms.AbsolutePath(dijkstrasPath)

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

func TestLargePathDijkstras(t *testing.T) {
	mazeSizeOriginal = mazeSizeOriginal * mazeSizeOriginal

	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(mazeSizeOriginal)

	var dijkstrasPath = algorithms.Dijkstras(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
	var absolutePathDijkstras, weightDijkstras = algorithms.AbsolutePath(dijkstrasPath)

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

func TestSmallPathAStar(t *testing.T) {
	mazeSizeOriginal = 4
	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(mazeSizeOriginal)

	var aStarPath = algorithms.AStar(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
	var absolutePathAStar, weightAStar = algorithms.AbsolutePath(aStarPath)

	if len(absolutePathAStar) == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(absolutePathAStar), "Path Length = 0")
	} else if weightAStar == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", weightAStar, "Number = 0 ")
	} else if len(absolutePathAStar) == mazeSizeOriginal {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(absolutePathAStar), "Lenght of Maze ")
	} else if weightAStar == mazeSizeOriginal {
		t.Errorf("Result was incorrect, got: %d, want: %s.", weightAStar, "Length of Maze ")
	}
}

func TestMediumPathAStar(t *testing.T) {
	mazeSizeOriginal = mazeSizeOriginal * 2
	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(mazeSizeOriginal)

	var aStarPath = algorithms.AStar(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
	var absolutePathAStar, weightAStar = algorithms.AbsolutePath(aStarPath)

	if len(absolutePathAStar) == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(absolutePathAStar), "Path Length = 0")
	} else if weightAStar == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", weightAStar, "Number = 0 ")
	} else if len(absolutePathAStar) == mazeSizeOriginal {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(absolutePathAStar), "Lenght of Maze ")
	} else if weightAStar == mazeSizeOriginal {
		t.Errorf("Result was incorrect, got: %d, want: %s.", weightAStar, "Length of Maze ")
	}
}

func TestLargePathAStar(t *testing.T) {
	mazeSizeOriginal = mazeSizeOriginal * mazeSizeOriginal

	var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(mazeSizeOriginal)

	var aStarPath = algorithms.AStar(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
	var absolutePathAStar, weightAStar = algorithms.AbsolutePath(aStarPath)

	if len(absolutePathAStar) == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(absolutePathAStar), "Path Length = 0")
	} else if weightAStar == 0 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", weightAStar, "Number = 0 ")
	} else if len(absolutePathAStar) == mazeSizeOriginal {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(absolutePathAStar), "Lenght of Maze ")
	} else if weightAStar == mazeSizeOriginal {
		t.Errorf("Result was incorrect, got: %d, want: %s.", weightAStar, "Length of Maze ")
	}
}

func BenchmarkAbsolutePathDijkstra(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(mazeSizeOriginal)

		var dijkstrasPath = algorithms.Dijkstras(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
		algorithms.AbsolutePath(dijkstrasPath)
	}
}

func BenchmarkAbsolutePathAStar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(mazeSizeOriginal)

		var aStarPath = algorithms.AStar(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
		algorithms.AbsolutePath(aStarPath)
	}
}

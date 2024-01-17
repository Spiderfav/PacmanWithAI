package testing

import (
	"testing"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/generation"
)

func TestCreateGraph(t *testing.T) {
	dfsGrid := algorithms.DFS(180, nil)
	graph := generation.MazeToGraph(dfsGrid, 20, 20, 20*180, 20*180)

	if graph[0].NodePosition.XCoordinate != 20 {
		t.Errorf("Result was incorrect, got start node X Coordinate: %f, want: %s.", graph[0].NodePosition.XCoordinate, "20")

	} else if graph[0].NodePosition.YCoordinate != 20 {
		t.Errorf("Result was incorrect, got start node Y Coordinate: %f, want: %s.", graph[0].NodePosition.XCoordinate, "20")

	} else if graph[len(graph)-1].NodePosition.XCoordinate != 3600 {
		t.Errorf("Result was incorrect, got end node X Coordinate: %f, want: %s.", graph[len(graph)-1].NodePosition.XCoordinate, "3,600")

	} else if graph[len(graph)-1].NodePosition.YCoordinate != 3600 {
		t.Errorf("Result was incorrect, got end node Y Coordinate: %f, want: %s.", graph[len(graph)-1].NodePosition.XCoordinate, "3,600")
	}
}

func BenchmarkMazeToGraph(b *testing.B) {
	dfsGrid := algorithms.DFS(180, nil)
	for i := 0; i < b.N; i++ {
		generation.MazeToGraph(dfsGrid, 20, 20, 20*180, 20*180)
	}
}

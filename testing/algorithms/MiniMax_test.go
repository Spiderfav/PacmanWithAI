package testing

import (
	"math"
	"testing"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

func BenchmarkMiniMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(8, nil, 20)

		enemyPosArr := []mazegrid.Position{gameGridDFS[0][0].NodePosition}

		ghostPosArr := []mazegrid.Position{gameGridDFS[len(gameGridDFS[0])-1][len(gameGridDFS[0])-1].NodePosition}

		params := algorithms.PruningParams{Alpha: math.Inf(1), Beta: math.Inf(-1)}

		for i := 0; i < b.N; i++ {
			algorithms.MiniMax(gameGridDFS, params, enemyPosArr, 0, ghostPosArr, mazegrid.GetPellotsPos(gameGridDFS), 10, true, false, 20)

		}
	}
}

func BenchmarkMiniMaxPruning(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(8, nil, 20)

		enemyPosArr := []mazegrid.Position{gameGridDFS[0][0].NodePosition}

		ghostPosArr := []mazegrid.Position{gameGridDFS[len(gameGridDFS[0])-1][len(gameGridDFS[0])-1].NodePosition}

		params := algorithms.PruningParams{Alpha: math.Inf(1), Beta: math.Inf(-1)}

		for i := 0; i < b.N; i++ {
			algorithms.MiniMax(gameGridDFS, params, enemyPosArr, 0, ghostPosArr, mazegrid.GetPellotsPos(gameGridDFS), 10, true, true, 20)

		}
	}
}

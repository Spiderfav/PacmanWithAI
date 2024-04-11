package testing

import (
	"testing"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

func BenchmarkExpectimax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(8, nil, 20)

		enemyPosArr := []mazegrid.Position{gameGridDFS[0][0].NodePosition}

		ghostPosArr := []mazegrid.Position{gameGridDFS[len(gameGridDFS[0])-1][len(gameGridDFS[0])-1].NodePosition}

		for i := 0; i < b.N; i++ {
			algorithms.Expectimax(gameGridDFS, enemyPosArr, 0, ghostPosArr, mazegrid.GetPellotsPos(gameGridDFS), 10, true, 20)
		}
	}
}

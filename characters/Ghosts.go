package characters

import (
	"context"
	"fmt"
	"image/color"
	_ "image/png"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

type NPC struct {
	Attributes Character
	Algo       algorithms.Algorithm
	Path       []mazegrid.MazeSquare
	hasMutex   bool
	Ctx        context.Context
	CancelFunc context.CancelFunc
	Pellots    []mazegrid.Position
	Cooldown   int
}

func (npc *NPC) Init(pos mazegrid.Position, colour color.Color, algo algorithms.Algorithm, enemyPos mazegrid.Position, grid [][]mazegrid.MazeSquare) {
	npc.Attributes.Init(pos, colour)
	npc.Algo = algo
	npc.Pellots = algorithms.GetPellotsPos(grid)
	npc.Path = npc.calculatePath(enemyPos, 0, grid)
	npc.hasMutex = true
	npc.Cooldown = 0
	npc.Ctx, npc.CancelFunc = context.WithCancel(context.Background())

}

func (npc *NPC) CancelContext() {
	if npc.CancelFunc != nil {
		npc.CancelFunc()
		npc.CancelFunc = nil
	}
}

func (npc *NPC) GetPosition() mazegrid.Position {
	return npc.Attributes.GetPosition()
}

func (npc *NPC) UpdatePosition(pos mazegrid.Position, enemyPos mazegrid.Position, enemyPoints int, grid [][]mazegrid.MazeSquare) {
	npc.Attributes.SetPosition(pos)

	if npc.Cooldown == 3 || len(npc.Path) < 2 {
		fmt.Println("Creating new path")
		npc.Pellots = algorithms.GetPellotsPos(grid)
		npc.Path = npc.calculatePath(enemyPos, enemyPoints, grid)

		npc.Cooldown = 0
	} else {
		fmt.Println("Taking already made path")
		npc.Path = npc.Path[:len(npc.Path)-1]
		npc.Cooldown += 1
	}

}

func (npc *NPC) GetAlgo() int {
	return npc.Algo
}

func (npc *NPC) calculatePath(enemyPos mazegrid.Position, enemyPoints int, grid [][]mazegrid.MazeSquare) []mazegrid.MazeSquare {
	var path []mazegrid.MazeSquare
	switch npc.Algo {
	case algorithms.DijkstraAlgo:
		path, _ = algorithms.AbsolutePath(algorithms.Dijkstras(grid, int(npc.Attributes.Position.YCoordinate), int(npc.Attributes.Position.XCoordinate), int(enemyPos.YCoordinate), int(enemyPos.XCoordinate)))

	case algorithms.AStarAlgo:
		path, _ = algorithms.AbsolutePath(algorithms.AStar(grid, int(npc.Attributes.Position.YCoordinate), int(npc.Attributes.Position.XCoordinate), int(enemyPos.YCoordinate), int(enemyPos.XCoordinate)))

	case algorithms.ReflexAlgo:
		path, _ = algorithms.AbsolutePath(algorithms.Reflex(grid, enemyPos, npc.Attributes.Position, npc.Pellots))

	case algorithms.MiniMaxAlgo:
		enemyPosArr := []mazegrid.Position{enemyPos}

		ghostPosArr := []mazegrid.Position{npc.Attributes.Position}

		params := algorithms.PruningParams{Alpha: math.Inf(1), Beta: math.Inf(-1)}

		_, _, ghostPosArrNew, _ := algorithms.MiniMax(grid, params, enemyPosArr, enemyPoints, ghostPosArr, npc.Pellots, 10, true, true)

		path = algorithms.ReversePath(algorithms.PosToNode(grid, ghostPosArrNew))
	}

	return path
}

func (npc *NPC) GetFrameProperties() FrameProperties {
	return npc.Attributes.GetFrameProperties()
}

func (npc *NPC) Move(enemyPos mazegrid.Position, enemyPoints int, grid [][]mazegrid.MazeSquare) {
	if npc.hasMutex {
		npc.hasMutex = false
		go npc.wait(enemyPos, enemyPoints, grid)

	}
}

func (npc *NPC) SetFrameProperties(fp FrameProperties) {
	npc.Attributes.SetFrameProperties(fp)
}

func (npc *NPC) UpdateCount() {
	npc.Attributes.Count += 1
}

func (npc *NPC) GetCount() int {
	return npc.Attributes.GetCount()
}

func (npc *NPC) GetSprite() *ebiten.Image {
	return npc.Attributes.GetSprite()
}

func (npc *NPC) wait(enemyPos mazegrid.Position, enemyPoints int, grid [][]mazegrid.MazeSquare) {
	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop()

	for {
		select {
		case <-npc.Ctx.Done():
			npc.hasMutex = true
			return // Exit the loop if context is cancelled
		case <-ticker.C:
			nextNode := len(npc.Path) - 2

			if nextNode < 0 {
				nextNode = 0
			}

			fmt.Println("Current path to take before update: ", justPositions(npc.Path))
			npc.UpdatePosition(npc.Path[nextNode].NodePosition, enemyPos, enemyPoints, grid)
			fmt.Println("Current path to take: ", justPositions(npc.Path))
			fmt.Println("")
			fmt.Println("")
			npc.hasMutex = true
			return
		}
	}

}

func justPositions(path []mazegrid.MazeSquare) []mazegrid.Position {
	var posArr []mazegrid.Position

	for i := 0; i < len(path); i++ {
		posArr = append(posArr, path[i].NodePosition)
	}

	return posArr
}

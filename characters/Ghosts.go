package characters

import (
	"context"
	"fmt"
	_ "image/png"
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
}

func (npc *NPC) Init(pos mazegrid.Position, algo algorithms.Algorithm, enemyPos mazegrid.Position, grid [][]mazegrid.MazeSquare) {
	npc.Attributes.Init(pos)
	npc.Algo = algo
	npc.Path = npc.calculatePath(pos, enemyPos, grid)
	npc.hasMutex = true

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

func (npc *NPC) UpdatePosition(pos mazegrid.Position, enemyPos mazegrid.Position, grid [][]mazegrid.MazeSquare) {
	//fmt.Println("Pos before:", npc.Attributes.Position)
	npc.Attributes.SetPosition(pos)
	//fmt.Println("Pos after:", npc.Attributes.Position)

	// fmt.Println("Path to take:", npc.Path)
	// fmt.Println("Pos of path to to take:", npc.Path[len(npc.Path)-2])
	npc.Path = npc.calculatePath(pos, enemyPos, grid)
}

func (npc *NPC) GetAlgo() int {
	return npc.Algo
}

func (npc *NPC) calculatePath(pos mazegrid.Position, enemyPos mazegrid.Position, grid [][]mazegrid.MazeSquare) []mazegrid.MazeSquare {
	var path []mazegrid.MazeSquare
	switch npc.Algo {
	case algorithms.DijkstraAlgo:
		path, _ = algorithms.AbsolutePath(algorithms.Dijkstras(grid, int(pos.YCoordinate), int(pos.XCoordinate), int(enemyPos.YCoordinate), int(enemyPos.XCoordinate)))

	case algorithms.AStarAlgo:
		path, _ = algorithms.AbsolutePath(algorithms.AStar(grid, int(pos.YCoordinate), int(pos.XCoordinate), int(enemyPos.YCoordinate), int(enemyPos.XCoordinate)))
	}

	return path
}

func (npc *NPC) GetFrameProperties() FrameProperties {
	return npc.Attributes.GetFrameProperties()
}

func (npc *NPC) Move(enemyPos mazegrid.Position, grid [][]mazegrid.MazeSquare) {
	if npc.hasMutex {
		npc.hasMutex = false
		go npc.wait(enemyPos, grid)

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

func (npc *NPC) wait(enemyPos mazegrid.Position, grid [][]mazegrid.MazeSquare) {
	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop()

	for {
		select {
		case <-npc.Ctx.Done():
			npc.hasMutex = true
			return // Exit the loop if context is cancelled
		case <-ticker.C:
			nextNode := len(npc.Path) - 2

			fmt.Println("Node value:", nextNode)

			if nextNode < 0 {
				nextNode = 0
			}

			fmt.Println("Node value After:", nextNode)

			npc.UpdatePosition(npc.Path[nextNode].NodePosition, enemyPos, grid)
			npc.hasMutex = true
			return
		}
	}

}

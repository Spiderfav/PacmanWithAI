package characters

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

type NPC struct {
	Attributes Character
	Algo       algorithms.Algorithm
	Path       []mazegrid.MazeSquare
}

func (npc *NPC) Init(pos mazegrid.Position, algo algorithms.Algorithm, enemyPos mazegrid.Position, grid [][]mazegrid.MazeSquare) {
	npc.Attributes.Init(pos)
	npc.Algo = algo
	npc.Path = npc.calculatePath(pos, enemyPos, grid)

}

func (npc *NPC) GetPosition() mazegrid.Position {
	return npc.Attributes.GetPosition()
}

func (npc *NPC) UpdatePosition(pos mazegrid.Position, enemyPos mazegrid.Position, grid [][]mazegrid.MazeSquare) {
	npc.Attributes.SetPosition(pos)
	npc.calculatePath(pos, enemyPos, grid)
}

func (npc *NPC) GetAlgo() int {
	return npc.Algo
}

func (npc *NPC) calculatePath(pos mazegrid.Position, enemyPos mazegrid.Position, grid [][]mazegrid.MazeSquare) []mazegrid.MazeSquare {
	var path []mazegrid.MazeSquare
	switch npc.Algo {
	case algorithms.DijkstraAlgo:
		path, _ = algorithms.AbsolutePath(algorithms.Dijkstras(grid, int(pos.XCoordinate), int(pos.YCoordinate), int(enemyPos.XCoordinate), int(enemyPos.YCoordinate)))

	case algorithms.AStarAlgo:
		path, _ = algorithms.AbsolutePath(algorithms.AStar(grid, int(pos.XCoordinate), int(pos.YCoordinate), int(enemyPos.XCoordinate), int(enemyPos.YCoordinate)))
	}

	return path
}

func (npc *NPC) GetFrameProperties() FrameProperties {
	return npc.Attributes.GetFrameProperties()
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

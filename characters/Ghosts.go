package characters

import (
	"context"
	"image/color"
	_ "image/png"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// The NPC Class is the class used for any AI ghosts (or pacman) to traverse the maze
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

// This function intialises the NPC variables and creates a starting path for the ghost to take
func (npc *NPC) Init(pos mazegrid.Position, colour color.Color, algo algorithms.Algorithm, enemyPos mazegrid.Position, grid [][]mazegrid.MazeSquare) {
	npc.Attributes.Init(pos, colour)
	npc.Algo = algo
	npc.Pellots = algorithms.GetPellotsPos(grid)
	npc.Path = npc.calculatePath(enemyPos, 0, grid)
	npc.hasMutex = true
	npc.Cooldown = 0
	npc.Ctx, npc.CancelFunc = context.WithCancel(context.Background())

}

// This function is called to cancel any functions running in parallel to the program
func (npc *NPC) CancelContext() {
	if npc.CancelFunc != nil {
		npc.CancelFunc()
		npc.CancelFunc = nil
	}
}

// Returns the position of the NPC
func (npc *NPC) GetPosition() mazegrid.Position {
	return npc.Attributes.GetPosition()
}

// This function will update the position of the ghost given a new position, the enemy's position and the game grid
func (npc *NPC) UpdatePosition(pos mazegrid.Position, enemyPos mazegrid.Position, enemyPoints int, grid [][]mazegrid.MazeSquare) {
	npc.Attributes.SetPosition(pos)

	// npc.Pellots = algorithms.GetPellotsPos(grid)
	// npc.Path = npc.calculatePath(enemyPos, enemyPoints, grid)

	// Makes sure that the NPC is not stuck just recalculating paths each time
	if npc.Cooldown == 3 || len(npc.Path) < 2 {
		npc.Pellots = algorithms.GetPellotsPos(grid)
		npc.Path = npc.calculatePath(enemyPos, enemyPoints, grid)

		npc.Cooldown = 0
	} else {
		npc.Path = npc.Path[:len(npc.Path)-1]
		npc.Cooldown += 1
	}

}

// Returns the algo the NPC is using for pathing
func (npc *NPC) GetAlgo() int {
	return npc.Algo
}

// This function calculates a path that the NPC should take given the enemy position, the enemy's points and the game grid
// It returns the path that NPC will take
func (npc *NPC) calculatePath(enemyPos mazegrid.Position, enemyPoints int, grid [][]mazegrid.MazeSquare) []mazegrid.MazeSquare {
	var path []mazegrid.MazeSquare
	switch npc.Algo {
	case algorithms.DijkstraAlgo:
		path, _ = algorithms.AbsolutePath(algorithms.Dijkstras(grid, int(npc.Attributes.Position.YCoordinate), int(npc.Attributes.Position.XCoordinate), int(enemyPos.YCoordinate), int(enemyPos.XCoordinate)))

	case algorithms.AStarAlgo:
		path, _ = algorithms.AbsolutePath(algorithms.AStar(grid, int(npc.Attributes.Position.YCoordinate), int(npc.Attributes.Position.XCoordinate), int(enemyPos.YCoordinate), int(enemyPos.XCoordinate), 20))

		//path = algorithms.AStar(grid, int(npc.Attributes.Position.YCoordinate), int(npc.Attributes.Position.XCoordinate), int(enemyPos.YCoordinate), int(enemyPos.XCoordinate), 20)

	case algorithms.ReflexAlgo:
		path, _ = algorithms.AbsolutePath(algorithms.Reflex(grid, enemyPos, npc.Attributes.Position, npc.Pellots, 20))

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

// This function allows the NPC to move in the game grid if it is currently allowed
// It will only move the NPC once every, 500 milliseconds
func (npc *NPC) Move(enemyPos mazegrid.Position, enemyPoints int, grid [][]mazegrid.MazeSquare) {
	if npc.hasMutex {
		// A mutex is used here as this function is called in the Update section of the game code and is called as much as possible
		// So to prevent the overwriting of the path for a ghost, a mutex must be used
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

// This function will make the NPC wait to move to the next position until the given time is up
func (npc *NPC) wait(enemyPos mazegrid.Position, enemyPoints int, grid [][]mazegrid.MazeSquare) {
	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop()

	for {
		select {

		// If the function was cancelled, break out and give up the path to be re-written
		case <-npc.Ctx.Done():
			npc.hasMutex = true
			return // Exit the loop if context is cancelled
		case <-ticker.C:
			nextNode := len(npc.Path) - 2

			if nextNode < 0 {
				nextNode = 0
			}

			npc.UpdatePosition(npc.Path[nextNode].NodePosition, enemyPos, enemyPoints, grid)
			npc.hasMutex = true
			return
		}
	}

}

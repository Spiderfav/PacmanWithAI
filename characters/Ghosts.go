package characters

import (
	"context"
	"image/color"
	_ "image/png"
	"math"
	"time"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// The NPC Class is the class used for any AI ghosts (or pacman) to traverse the maze
type NPC struct {
	Attributes     Character
	Algo           algorithms.Algorithm
	Path           []mazegrid.MazeSquare
	hasMutex       bool
	Ctx            context.Context
	CancelFunc     context.CancelFunc
	pellots        []mazegrid.Position
	speed          int
	cooldown       int
	mazeSquareSize int
}

// This function intialises the NPC variables and creates a starting path for the ghost to take
func (npc *NPC) Init(pos mazegrid.Position, colour color.Color, algo algorithms.Algorithm, enemyPos mazegrid.Position, grid [][]mazegrid.MazeSquare, pellots []mazegrid.Position, squareSize int) {
	npc.Attributes.Init(pos, colour)
	npc.Algo = algo
	npc.pellots = pellots
	npc.mazeSquareSize = squareSize
	npc.Path = npc.calculatePath(enemyPos, 0, grid)
	npc.hasMutex = true
	npc.cooldown = 0
	npc.speed = 500
	npc.Ctx, npc.CancelFunc = context.WithCancel(context.Background())

}

// This function is called to cancel any functions running in parallel to the program
func (npc *NPC) CancelContext() {
	if npc.CancelFunc != nil {
		npc.CancelFunc()
		npc.CancelFunc = nil
	}
}

// Increases the ghost's speed
func (npc *NPC) IncreaseSpeed() {
	npc.speed = npc.speed - 50
}

// Resets the ghost's speed to original value
func (npc *NPC) ResetSpeed() {
	npc.speed = 500
}

// Returns the position of the NPC
func (npc *NPC) GetPosition() mazegrid.Position {
	return npc.Attributes.GetPosition()
}

// This function will update the position of the ghost given a new position, the enemy's position and the game grid
func (npc *NPC) UpdatePosition(pos mazegrid.Position, enemyPos mazegrid.Position, enemyPoints int, grid [][]mazegrid.MazeSquare) {
	npc.Attributes.SetPosition(pos)

	// If there are not enough pellots in the maze
	if len(npc.pellots) < 2*len(grid[0]) {
		npc.cooldown = 3
	}

	// Makes sure that the NPC is not stuck just recalculating paths each time
	if npc.cooldown == 3 || len(npc.Path) < 2 {
		npc.pellots = mazegrid.GetPellotsPos(grid)
		npc.Path = npc.calculatePath(enemyPos, enemyPoints, grid)

		npc.cooldown = 0
	} else {
		npc.Path = npc.Path[:len(npc.Path)-1]
		npc.cooldown += 1
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

	if npc.Algo == algorithms.MiniMaxAlgo {

		enemyPosArr := []mazegrid.Position{enemyPos}

		ghostPosArr := []mazegrid.Position{npc.Attributes.Position}

		params := algorithms.PruningParams{Alpha: math.Inf(1), Beta: math.Inf(-1)}

		_, _, ghostPosArrNew, _ := algorithms.MiniMax(grid, params, enemyPosArr, enemyPoints, ghostPosArr, npc.pellots, 10, true, true, npc.mazeSquareSize)

		path = algorithms.ReversePath(algorithms.PosToNode(grid, ghostPosArrNew, npc.mazeSquareSize))

	} else if npc.Algo == algorithms.ExpectimaxAlgo {

		enemyPosArr := []mazegrid.Position{enemyPos}

		ghostPosArr := []mazegrid.Position{npc.Attributes.Position}

		_, _, ghostPosArrNew := algorithms.Expectimax(grid, enemyPosArr, enemyPoints, ghostPosArr, npc.pellots, 10, true, npc.mazeSquareSize)

		path = algorithms.ReversePath(algorithms.PosToNode(grid, ghostPosArrNew, npc.mazeSquareSize))

	} else {

		path = algorithms.Reflex(grid, enemyPos, npc.Attributes.Position, npc.pellots, npc.mazeSquareSize, npc.Algo)

	}

	return path
}

// This function allows the NPC to move in the game grid if it is currently allowed
// It will only move the NPC once every, 500 milliseconds
func (npc *NPC) Move(enemyPos mazegrid.Position, enemyPoints int, grid [][]mazegrid.MazeSquare) {
	if npc.hasMutex {
		// A mutex is used here as this function is called in the Update section of the game code and is called as much as possible
		// So to prevent the overwriting of the path for a ghost, a mutex must be used
		npc.hasMutex = false
		npc.wait(enemyPos, enemyPoints, grid)

	}
}

func (npc *NPC) ResetMutex() {
	npc.hasMutex = true
}

// This function will make the NPC wait to move to the next position until the given time is up
func (npc *NPC) wait(enemyPos mazegrid.Position, enemyPoints int, grid [][]mazegrid.MazeSquare) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(npc.speed))
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

// This function, given the array of ghosts, changes their algorithm
func ChangeGhostsAlgo(ghosts []*NPC, ghostNewAlgo int) {

	for i := range ghosts {
		ghosts[i].Algo = ghostNewAlgo
	}

}

// This functions, given the array of ghosts, the game grid and the player, resets the ghost's path and updates their position
func ResetMovement(ghosts []*NPC, game mazegrid.Maze, player *Player) {

	newPath := []mazegrid.MazeSquare{game.Grid[game.Size/2][game.Size/2], game.Grid[game.Size/2][game.Size/2], game.Grid[game.Size/2][game.Size/2]}

	for i := range ghosts {
		if ghosts[i].CancelFunc != nil {
			ghosts[i].CancelFunc()
		}

		// Cancel any ghosts undergoing movement
		ghosts[i].Ctx, ghosts[i].CancelFunc = context.WithCancel(context.Background())

		ghosts[i].UpdatePosition(game.Grid[game.Size/2][game.Size/2].NodePosition, player.GetPosition(), 0, game.Grid)
		ghosts[i].Path = newPath

	}
}

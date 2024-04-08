package main

import (
	"context"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/characters"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// This function, given the array of ghosts, changes their algorithm
func changeGhostsAlgo(ghosts []*characters.NPC, ghostNewAlgo int) {

	for i := range ghosts {
		ghosts[i].Algo = ghostNewAlgo
	}

}

// This functions, given the array of ghosts, the game grid and the player, resets the ghost's path and updates their position
func update(ghosts []*characters.NPC, game mazegrid.Maze, player *characters.Player) {

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

// This function is called every game update to move the ghosts.
// It checks the ghost's and player position and pellots to end the game or just moves the ghosts in their path
func moveGhosts(g *Game) {

	// Game Over or new game
	for i := range g.Ghosts {
		if g.Ghosts[i].GetPosition() == g.Player.GetPosition() || len(g.Maze.Pellots) == 0 {
			changeMazeSize(g.Maze.Size, false, g)
			break
		}

		// Move the ghosts
		g.Ghosts[i].Move(g.Player.GetPosition(), g.Player.GetPoints(), g.Maze.Grid)

	}
}

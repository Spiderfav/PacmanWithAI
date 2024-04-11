package main

import (
	"image/color"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/characters"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/file"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/input"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

// Checks whether or not to draw the paths the ghosts will take
var drawGhostLines bool

// Checks if the game is menu or playing the game
var menuOrGame = 0

// The original grid size
const mazeSizeOriginal = 8

// Sets the size the maze will be rendered
const (
	squareSize = 30
	halfSquare = squareSize / 2
)

type GameButtons struct {
	buttonsSize  []*input.Button
	buttonsAlgo  []*input.Button
	buttonsGhost []*input.Button
	buttonBack   *input.Button
}

// Defines what a game object should keep track of
type Game struct {
	Maze        mazegrid.Maze
	buttonsMenu []*input.Button
	GameButtons
	fontFace font.Face
	Ghosts   []*characters.NPC
	Player   *characters.Player
}

var lastPlayerPoints int

// This function updates the game logic 60 times a second
func (g *Game) Update() error {

	if menuOrGame == 1 {

		g.Maze.Pellots = mazegrid.GetPellotsPos(g.Maze.Grid)

		if len(g.Maze.Pellots) == 0 {
			g.Player.ResetMapPoints()
			g.changeMazeSize(g.Maze.Size, false)

		}

		if g.Player.GetTotalPoints()%50 == 0 && lastPlayerPoints != g.Player.GetTotalPoints() {
			//Increase ghost speed
			lastPlayerPoints = g.Player.GetTotalPoints()
			g.increaseGhostSpeed()
		}

		// Check if player is moving
		go g.Player.IsPlayerMoving(g.Maze.Grid, squareSize)

		// Move ghosts
		go g.moveGhosts()

	}

	// Check if the mouse button is clicked on a given button
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		// If the menu buttons are enabled
		if g.buttonsMenu[0].Enabled {
			if g.buttonsMenu[0].In(x, y) {
				menuOrGame = 1
				g.buttonBack.Enabled = true

				// Change Size and Algo buttons to enabled and disable the Menu Buttons
				g.changeStates(true)
				return nil

			} else if g.buttonsMenu[1].In(x, y) {

				g.Maze.Grid = file.LoadFromFile()
				g.Maze.Size = len(g.Maze.Grid[0])
				g.changeMazeSize(0, true)

			} else if g.buttonsMenu[2].In(x, y) {
				drawGhostLines = !drawGhostLines

			} else if g.buttonsMenu[3].In(x, y) {
				os.Exit(0)
			}

			// If the game buttons are enabled
		} else if g.buttonBack.Enabled {

			if g.buttonBack.In(x, y) {
				menuOrGame = 0
				g.buttonBack.Enabled = false

				// Change Size and Algo buttons to disabled and enable the Menu Buttons
				g.changeStates(false)
				return nil

			} else if g.buttonsSize[0].In(x, y) {
				input.ResetColours(g.buttonsSize)
				g.buttonsSize[0].ChangeColour(color.RGBA{0, 255, 0, 250})
				g.changeMazeSize(mazeSizeOriginal, false)

			} else if g.buttonsSize[1].In(x, y) {
				input.ResetColours(g.buttonsSize)
				g.buttonsSize[1].ChangeColour(color.RGBA{0, 255, 0, 250})
				g.changeMazeSize(mazeSizeOriginal*2, false)

			} else if g.buttonsSize[2].In(x, y) {
				input.ResetColours(g.buttonsSize)
				g.buttonsSize[2].ChangeColour(color.RGBA{0, 255, 0, 250})
				g.changeMazeSize(mazeSizeOriginal*3, false)

			} else if g.buttonsSize[3].In(x, y) {
				file.SaveToFile(g.Maze.Grid)

			} else if g.buttonsAlgo[0].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[0].ChangeColour(color.RGBA{0, 255, 0, 250})
				characters.ChangeGhostsAlgo(g.Ghosts, algorithms.DijkstraAlgo)

			} else if g.buttonsAlgo[1].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[1].ChangeColour(color.RGBA{0, 255, 0, 250})
				characters.ChangeGhostsAlgo(g.Ghosts, algorithms.AStarAlgo)

			} else if g.buttonsAlgo[2].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[2].ChangeColour(color.RGBA{0, 255, 0, 250})
				characters.ChangeGhostsAlgo(g.Ghosts, algorithms.BFSAlgo)

			} else if g.buttonsAlgo[3].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[3].ChangeColour(color.RGBA{0, 255, 0, 250})
				characters.ChangeGhostsAlgo(g.Ghosts, algorithms.DFSAlgo)

			} else if g.buttonsAlgo[4].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[4].ChangeColour(color.RGBA{0, 255, 0, 250})
				characters.ChangeGhostsAlgo(g.Ghosts, algorithms.MiniMaxAlgo)

			} else if g.buttonsAlgo[5].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[5].ChangeColour(color.RGBA{0, 255, 0, 250})
				characters.ChangeGhostsAlgo(g.Ghosts, algorithms.ExpectimaxAlgo)

			} else if g.buttonsGhost[0].In(x, y) {
				characters.ResetMovement(g.Ghosts, g.Maze, g.Player)
				ghostNew := &characters.NPC{}
				ghostNew.Init(g.Maze.Grid[g.Maze.Size/2][g.Maze.Size/2].NodePosition, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}, g.Ghosts[0].Algo, g.Player.GetPosition(), g.Maze.Grid, g.Maze.Pellots, squareSize)
				g.Ghosts = append(g.Ghosts, ghostNew)

			} else if g.buttonsGhost[1].In(x, y) {
				if len(g.Ghosts) > 1 {
					g.Ghosts = g.Ghosts[:len(g.Ghosts)-1]
					characters.ResetMovement(g.Ghosts, g.Maze, g.Player)
				}

			}

		}
	}

	return nil
}

// This function is called to change the buttons shown on screen
func (g *Game) changeStates(shown bool) {
	input.ChangeStateButtons(g.buttonsSize[:], shown)
	input.ChangeStateButtons(g.buttonsAlgo[:], shown)
	input.ChangeStateButtons(g.buttonsGhost[:], shown)
	input.ChangeStateButtons(g.buttonsMenu[:], !shown)
}

// This function is called every second to update what is drawn on the screen
func (g *Game) Draw(screen *ebiten.Image) {

	switch menuOrGame {
	case 0:
		mainMenu(screen, g)

	case 1:
		gameMenu(screen, g)
		DrawSprite(screen, g.Player.Attributes)

		for _, ghosts := range g.Ghosts {
			DrawSprite(screen, ghosts.Attributes)

			if drawGhostLines {
				drawPathsLines(screen, ghosts.Path)

			}
		}

	}

}

// This function dictates the size of the window for the game
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Takes 3 parameters: The new maze size, a boolean flag for if we are loading a maze from the file and the game object
// This function, will check if we are loading the maze and if so, loads the given maze or, updates the game to the new maze size
func (g *Game) changeMazeSize(newSize int, loadedMaze bool) {

	// If maze isn't being loaded
	if !loadedMaze {
		// Create new maze with the given size
		g.Maze = algorithms.CreateMaze(newSize, squareSize)

		g.Player.ResetMapPoints()

	}

	// Reset the player and the Ghosts position to their original start
	g.Player.SetPosition(g.Maze.Grid[0][0].NodePosition)

	characters.ResetMovement(g.Ghosts, g.Maze, g.Player)

}

// This function is called every game update to move the ghosts.
// It checks the ghost's and player position and pellots to end the game or just moves the ghosts in their path
func (g *Game) moveGhosts() {

	// Checking the position of every ghost
	for i := range g.Ghosts {

		// If the ghost has caught the player
		if g.Ghosts[i].GetPosition() == g.Player.GetPosition() {

			// Remove a life from the player
			g.Player.RemoveLife()

			// If the player's lives are zero, reset the maze
			if g.Player.GetLives() == -1 {
				g.changeMazeSize(g.Maze.Size, false)
				g.Player.GameOver()
				g.resetGhostSpeed()
				break
			}

			// As player was caught, restart the map
			g.changeMazeSize(g.Maze.Size, true)
			break
		}

		// Move the ghosts
		g.Ghosts[i].Move(g.Player.GetPosition(), g.Player.GetMapPoints(), g.Maze.Grid)

	}
}

// This function resets the ghost's speed back to their original when Pacman loses
func (g *Game) resetGhostSpeed() {
	for _, ghosts := range g.Ghosts {
		ghosts.ResetSpeed()
	}
}

// This function increases the ghost's speed based on how many points Pacman has
func (g *Game) increaseGhostSpeed() {
	for _, ghosts := range g.Ghosts {
		ghosts.IncreaseSpeed()
	}
}

// This function is called by the game object to create the game environment
func NewGame() *Game {

	maze := algorithms.CreateMaze(mazeSizeOriginal, squareSize)

	// Creating the player object
	pacman := characters.Player{}
	pacman.Init(maze.Grid[0][0].NodePosition, color.RGBA{255, 234, 0, 255}, 3)

	// Creating the Enemy
	ghost := characters.NPC{}
	ghost.Init(maze.Grid[mazeSizeOriginal/2][mazeSizeOriginal/2].NodePosition, color.RGBA{200, 0, 0, 255}, algorithms.AStarAlgo, pacman.GetPosition(), maze.Grid, maze.Pellots, squareSize)
	ghosts := []*characters.NPC{&ghost}

	// Initialize all buttons
	buttonImage := ebiten.NewImage(100, 30)        // Set the size of the button
	buttonImage.Fill(color.RGBA{0, 255, 255, 250}) // Fill with a color

	buttonsMenu := input.MakeMainMenuButtons(screenWidth, screenHeight)

	buttonsSize := input.GameSizeButtons(screenWidth, screenHeight)

	buttonsAlgo := input.GameAlgoButtons(screenWidth, screenHeight)

	buttonsGhost := input.GameGhostButtons(screenWidth, screenHeight)

	buttonBack := &input.Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 50, // Position of the button
		Y:       (screenHeight / 2),
		Width:   100,
		Height:  30,
		Message: "Main Menu",
		Enabled: false,
	}

	// Initialize the game.
	return &Game{
		buttonsMenu: buttonsMenu,
		GameButtons: GameButtons{buttonsSize, buttonsAlgo, buttonsGhost, buttonBack},
		fontFace:    basicfont.Face7x13,
		Ghosts:      ghosts,
		Maze:        maze,
		Player:      &pacman,
	}
}

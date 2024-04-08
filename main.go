package main

import (
	"image/color"
	"log"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kbinani/screenshot"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/characters"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/file"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/input"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

// These variables set the game size to the screen the user is running it on
var (
	screenWidth  = screenshot.GetDisplayBounds(0).Dx()
	screenHeight = screenshot.GetDisplayBounds(0).Dy()
)

// Sets the size the maze will be rendered
const (
	squareSize = 30
	halfSquare = squareSize / 2
)

// Defines what a game object should keep track of
type Game struct {
	Maze         mazegrid.Maze
	buttonsMenu  []*input.Button
	buttonsSize  []*input.Button
	buttonsAlgo  []*input.Button
	buttonsGhost []*input.Button
	buttonBack   *input.Button
	fontFace     font.Face
	Ghosts       []*characters.NPC
	Player       *characters.Player
}

// The original grid size
const mazeSizeOriginal = 8

// Checks if the game is menu or playing the game
var menuOrGame = 0

// This function updates the game logic 60 times a second
func (g *Game) Update() error {

	if menuOrGame == 1 {

		g.Maze.Pellots = mazegrid.GetPellotsPos(g.Maze.Grid)

		// Check if player is moving
		go g.Player.IsPlayerMoving(g.Maze.Grid, squareSize)

		// go moveGhosts(g) // Causes memory leak
		moveGhosts(g)

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
				input.ChangeStateButtons(g.buttonsSize[:], true)
				input.ChangeStateButtons(g.buttonsAlgo[:], true)
				input.ChangeStateButtons(g.buttonsGhost[:], true)
				input.ChangeStateButtons(g.buttonsMenu[:], false)
				return nil

			} else if g.buttonsMenu[1].In(x, y) {
				g.Maze.Grid = file.LoadFromFile()
				g.Maze.Size = len(g.Maze.Grid[0])
				changeMazeSize(0, true, g)
			} else if g.buttonsMenu[2].In(x, y) {

			} else if g.buttonsMenu[3].In(x, y) {
				os.Exit(0)
			}

			// If the game buttons are enabled
		} else if g.buttonBack.Enabled {

			if g.buttonBack.In(x, y) {
				menuOrGame = 0
				g.buttonBack.Enabled = false

				// Change Size and Algo buttons to disabled and enable the Menu Buttons
				input.ChangeStateButtons(g.buttonsSize[:], false)
				input.ChangeStateButtons(g.buttonsAlgo[:], false)
				input.ChangeStateButtons(g.buttonsGhost[:], false)
				input.ChangeStateButtons(g.buttonsMenu[:], true)
				return nil

			} else if g.buttonsSize[0].In(x, y) {
				input.ResetColours(g.buttonsSize)
				g.buttonsSize[0].ChangeColour(color.RGBA{0, 255, 0, 250})
				changeMazeSize(mazeSizeOriginal, false, g)

			} else if g.buttonsSize[1].In(x, y) {
				input.ResetColours(g.buttonsSize)
				g.buttonsSize[1].ChangeColour(color.RGBA{0, 255, 0, 250})
				changeMazeSize(mazeSizeOriginal*2, false, g)

			} else if g.buttonsSize[2].In(x, y) {
				input.ResetColours(g.buttonsSize)
				g.buttonsSize[2].ChangeColour(color.RGBA{0, 255, 0, 250})
				changeMazeSize((mazeSizeOriginal*2)*2, false, g)

			} else if g.buttonsSize[3].In(x, y) {
				file.SaveToFile(g.Maze.Grid)

			} else if g.buttonsAlgo[0].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[0].ChangeColour(color.RGBA{0, 255, 0, 250})
				changeGhostsAlgo(g.Ghosts, algorithms.DijkstraAlgo)

			} else if g.buttonsAlgo[1].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[1].ChangeColour(color.RGBA{0, 255, 0, 250})
				changeGhostsAlgo(g.Ghosts, algorithms.AStarAlgo)

			} else if g.buttonsAlgo[2].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[2].ChangeColour(color.RGBA{0, 255, 0, 250})
				changeGhostsAlgo(g.Ghosts, algorithms.BFSAlgo)

			} else if g.buttonsAlgo[3].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[3].ChangeColour(color.RGBA{0, 255, 0, 250})
				changeGhostsAlgo(g.Ghosts, algorithms.DFSAlgo)

			} else if g.buttonsAlgo[4].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[4].ChangeColour(color.RGBA{0, 255, 0, 250})
				changeGhostsAlgo(g.Ghosts, algorithms.MiniMaxAlgo)

			} else if g.buttonsGhost[0].In(x, y) {
				update(g.Ghosts, g.Maze, g.Player)
				ghostNew := &characters.NPC{}
				ghostNew.Init(g.Maze.Grid[g.Maze.Size/2][g.Maze.Size/2].NodePosition, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}, g.Ghosts[0].Algo, g.Player.GetPosition(), g.Maze.Grid, g.Maze.Pellots, squareSize)
				g.Ghosts = append(g.Ghosts, ghostNew)

			} else if g.buttonsGhost[1].In(x, y) {
				if len(g.Ghosts) > 1 {
					g.Ghosts = g.Ghosts[:len(g.Ghosts)-1]
					update(g.Ghosts, g.Maze, g.Player)
				}

			}

		}
	}

	return nil
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
			drawPathsLines(screen, ghosts.Path)
		}

	}

}

// This function dictates the size of the window for the game
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// This function is called by the game object to create the game environment
func NewGame() *Game {

	// Creating the maze by aplying DFS twice
	oldGameGridDFS := algorithms.DFS(mazeSizeOriginal, nil, squareSize)
	algorithms.MarkUnvisited(oldGameGridDFS, false)
	gameGridDFS := algorithms.DFS(mazeSizeOriginal, oldGameGridDFS, squareSize)
	maze := mazegrid.Maze{Size: mazeSizeOriginal, Grid: gameGridDFS, Pellots: mazegrid.GetPellotsPos(gameGridDFS)}

	// Creating the player object
	pacman := characters.Player{}
	pacman.Init(gameGridDFS[0][0].NodePosition, color.RGBA{255, 234, 0, 255})

	// Creating the Enemy
	ghost := characters.NPC{}
	ghost.Init(gameGridDFS[mazeSizeOriginal/2][mazeSizeOriginal/2].NodePosition, color.RGBA{200, 0, 0, 255}, algorithms.DFSAlgo, pacman.GetPosition(), gameGridDFS, maze.Pellots, squareSize)
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
		buttonsMenu:  buttonsMenu,
		buttonsSize:  buttonsSize,
		buttonsAlgo:  buttonsAlgo,
		buttonsGhost: buttonsGhost,
		buttonBack:   buttonBack,
		fontFace:     basicfont.Face7x13,
		Ghosts:       ghosts,
		Maze:         maze,
		Player:       &pacman,
	}
}

// Takes 3 parameters: The new maze size, a boolean flag for if we are loading a maze from the file and the game object
// This function, will check if we are loading the maze and if so, loads the given maze or, updates the game to the new maze size
func changeMazeSize(newSize int, loadedMaze bool, g *Game) {

	// If maze isn't being loaded
	if !loadedMaze {
		// Create new maze with the given size
		oldGameGridDFS := algorithms.DFS(newSize, nil, squareSize)
		algorithms.MarkUnvisited(oldGameGridDFS, false)
		g.Maze.Grid = algorithms.DFS(newSize, oldGameGridDFS, squareSize)
		// Set new game size to be the given size
		g.Maze.Size = newSize

	}

	// Reset the player and the Ghosts position to their original start
	g.Player.SetPosition(g.Maze.Grid[0][0].NodePosition)
	g.Player.ResetPoints()

	update(g.Ghosts, g.Maze, g.Player)

}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pacman")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

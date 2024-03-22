package main

import (
	"context"
	"image/color"
	"log"
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
	Maze        mazegrid.Maze
	buttonsMenu []*input.Button
	buttonsSize []*input.Button
	buttonsAlgo []*input.Button
	buttonBack  *input.Button
	fontFace    font.Face
	Ghosts      characters.NPC
	Player      characters.Player
}

// The original grid size
const mazeSizeOriginal = 8

// Checks if the game is menu or playing the game
var menuOrGame = 0

// This function updates the game logic 60 times a second
func (g *Game) Update() error {
	g.Ghosts.UpdateCount()
	g.Maze.Pellots = mazegrid.GetPellotsPos(g.Maze.Grid)

	if menuOrGame == 1 {
		// Game Over or new game
		if g.Ghosts.GetPosition() == g.Player.GetPosition() || len(g.Maze.Pellots) == 0 {
			changeMazeSize(g.Maze.Size, false, g)
		}

		// Check if player is moving
		go g.Player.IsPlayerMoving(g.Maze.Grid, squareSize)

		// Move the ghosts
		g.Ghosts.Move(g.Player.GetPosition(), g.Player.GetPoints(), g.Maze.Grid)

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
				g.Ghosts.Algo = algorithms.DijkstraAlgo

			} else if g.buttonsAlgo[1].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[1].ChangeColour(color.RGBA{0, 255, 0, 250})
				g.Ghosts.Algo = algorithms.AStarAlgo

			} else if g.buttonsAlgo[2].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[2].ChangeColour(color.RGBA{0, 255, 0, 250})
				g.Ghosts.Algo = algorithms.BFSAlgo

			} else if g.buttonsAlgo[3].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[3].ChangeColour(color.RGBA{0, 255, 0, 250})
				g.Ghosts.Algo = algorithms.DFSAlgo

			} else if g.buttonsAlgo[4].In(x, y) {
				input.ResetColours(g.buttonsAlgo)
				g.buttonsAlgo[4].ChangeColour(color.RGBA{0, 255, 0, 250})
				g.Ghosts.Algo = algorithms.MiniMaxAlgo

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
		DrawSprite(screen, g.Ghosts.Attributes)
		DrawSprite(screen, g.Player.Attributes)
		drawPathsLines(screen, g.Ghosts.Path)

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
	ghost.Init(gameGridDFS[mazeSizeOriginal/2][mazeSizeOriginal/2].NodePosition, color.RGBA{200, 0, 0, 255}, algorithms.ReflexAlgo, pacman.GetPosition(), gameGridDFS, maze.Pellots, squareSize)

	// Initialize all buttons
	buttonImage := ebiten.NewImage(100, 30)        // Set the size of the button
	buttonImage.Fill(color.RGBA{0, 255, 255, 250}) // Fill with a color

	buttonsMenu := input.MakeMainMenuButtons(screenWidth, screenHeight)

	buttonsSize := input.GameSizeButtons(screenWidth, screenHeight)

	buttonsAlgo := input.GameAlgoButtons(screenWidth, screenHeight)

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
		buttonsSize: buttonsSize,
		buttonsAlgo: buttonsAlgo,
		buttonBack:  buttonBack,
		fontFace:    basicfont.Face7x13,
		Ghosts:      ghost,
		Maze:        maze,
		Player:      pacman,
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

	if g.Ghosts.CancelFunc != nil {
		g.Ghosts.CancelFunc()
	}

	// Cancel any ghosts undergoing movement
	g.Ghosts.Ctx, g.Ghosts.CancelFunc = context.WithCancel(context.Background())

	// Reset the player and the Ghosts position to their original start
	g.Player.SetPosition(g.Maze.Grid[0][0].NodePosition)
	g.Player.ResetPoints()

	g.Ghosts.UpdatePosition(g.Maze.Grid[g.Maze.Size/2][g.Maze.Size/2].NodePosition, g.Player.GetPosition(), 0, g.Maze.Grid)

}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pacman")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

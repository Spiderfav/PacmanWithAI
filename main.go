package main

import (
	"context"
	"image/color"
	"log"

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

const (
	screenWidth  = 1920
	screenHeight = 1080
	squareSize   = 30
	halfSquare   = squareSize / 2
)

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

const mazeSizeOriginal = 8

var menuOrGame = 0

// This function updated the game logic 60 times a second
func (g *Game) Update() error {
	g.Ghosts.UpdateCount()
	g.Maze.Pellots = mazegrid.GetPellotsPos(g.Maze.Grid)

	if menuOrGame == 1 {
		// Checking if the player is moving and if so, moving the player
		if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {

			g.Player.Move(characters.Up, g.Maze.Grid, squareSize)

		} else if inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {

			g.Player.Move(characters.Down, g.Maze.Grid, squareSize)

		} else if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
			g.Player.Move(characters.Left, g.Maze.Grid, squareSize)

		} else if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
			g.Player.Move(characters.Right, g.Maze.Grid, squareSize)

		}

		g.Ghosts.Move(g.Player.GetPosition(), g.Player.GetPoints(), g.Maze.Grid)

		// Game Over or new game
		if g.Ghosts.GetPosition() == g.Player.GetPosition() || len(g.Maze.Pellots) == 0 {
			changeMazeSize(g.Maze.Size, false, g)
		}
	}

	// Check if the mouse button is clicked on a given button
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if g.buttonsMenu[0].Enabled {
			if g.buttonsMenu[0].In(x, y) {
				menuOrGame = 1
				g.buttonBack.Enabled = true
				input.ChangeStateButtons(g.buttonsSize[:], true)
				input.ChangeStateButtons(g.buttonsAlgo[:], true)
				input.ChangeStateButtons(g.buttonsMenu[:], false)
				return nil

			} else if g.buttonsMenu[1].In(x, y) {
				g.Maze.Grid = file.LoadFromFile()
				g.Maze.Size = len(g.Maze.Grid[0])
				changeMazeSize(0, true, g)
			}

		} else if g.buttonBack.Enabled {

			if g.buttonBack.In(x, y) {
				menuOrGame = 0
				g.buttonBack.Enabled = false
				input.ChangeStateButtons(g.buttonsSize[:], false)
				input.ChangeStateButtons(g.buttonsAlgo[:], false)
				input.ChangeStateButtons(g.buttonsMenu[:], true)
				return nil

			} else if g.buttonsSize[0].In(x, y) {
				changeMazeSize(mazeSizeOriginal, false, g)

			} else if g.buttonsSize[1].In(x, y) {
				changeMazeSize(mazeSizeOriginal*2, false, g)

			} else if g.buttonsSize[2].In(x, y) {
				changeMazeSize((mazeSizeOriginal*2)*2, false, g)

			} else if g.buttonsSize[3].In(x, y) {
				file.SaveToFile(g.Maze.Grid)

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
	ghost.Init(gameGridDFS[mazeSizeOriginal/2][mazeSizeOriginal/2].NodePosition, color.RGBA{200, 0, 0, 255}, algorithms.AStarAlgo, pacman.GetPosition(), gameGridDFS, maze.Pellots, squareSize)

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
